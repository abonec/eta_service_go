package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v3"
	"math"
	"strconv"
)

type DbQuery struct {
	index       string
	indexType   string
	lat         float64
	lon         float64
	distance    string
	etaModifier float64
	cabSize     int
	client      *elastic.Client
}

func NewDbQuery(index string) *DbQuery {
	return &DbQuery{
		index:       index,
		indexType:   "cab",
		distance:    "30km",
		etaModifier: 1.5,
		cabSize:     3,
		client:      database,
	}
}

func (finder *DbQuery) GetEta(lat, lon float64, vacant bool) float64 {
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("vacant", vacant))
	distanceQuery := elastic.NewGeoDistanceQuery("location").Distance(finder.distance).Point(lat, lon)
	query = query.Filter(distanceQuery)
	sorter := elastic.NewGeoDistanceSort("location").Point(lat, lon).Unit("km").GeoDistance("sloppy_arc").Asc()
	result, err := finder.client.Search().Index(finder.index).Type(finder.indexType).
		Query(query).SortBy(sorter).Size(finder.cabSize).Do()
	HandleError(err)
	distance_sum := 0.0
	if result.Hits != nil {
		for _, hit := range result.Hits.Hits {
			distance_sum += hit.Sort[0].(float64)
		}
	}
	eta := distance_sum / float64(len(result.Hits.Hits)) * finder.etaModifier
	if math.IsNaN(eta) {
		return -1
	}
	return eta
}

func (finder *DbQuery) BulkIndex(cabs []*Cab, async bool, addId bool) {
	bulk := finder.client.Bulk().Refresh(!async)
	for i := 0; i < len(cabs); i++ {
		request := elastic.NewBulkIndexRequest().Index(finder.index).Type(finder.indexType).Doc(cabs[i])
		if addId {
			request = request.Id(strconv.Itoa(i))
		}
		bulk.Add(request)
	}

	_, err := bulk.Do()
	HandleError(err)
}

func (finder *DbQuery) IndexSize() int64 {
	count, err := finder.client.Count(finder.index).Do()
	HandleError(err)
	return count
}

func (finder *DbQuery) GetById(id int) *Cab {
	get, err := finder.client.Get().Index(finder.index).Type(finder.indexType).Id(strconv.Itoa(id)).Do()
	HandleError(err)
	if get.Found {
		return NewCabFromJson(*get.Source)
	}
	return nil
}

func (finder *DbQuery) Put(cab *Cab) {
	response, err := finder.client.Index().Index(finder.index).Type(finder.indexType).BodyJson(cab).Do()
	failOnError(err, "failed to add a cab")
	broadcast := fmt.Sprintf("%s with id '%s' was sent to index '%s' and has created status as %t",
		response.Type, response.Id, response.Index, response.Created)
	amqpCabUpdatedExchange.Publish([]byte(broadcast))
}
