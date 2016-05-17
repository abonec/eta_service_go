package main

import "gopkg.in/olivere/elastic.v3"

type DbQuery struct {
	index string
	indexType string
	lat float64
	lon float64
	distance string
	etaModifier float64
	cabSize int
}

func NewDbQuery() *DbQuery {
	return &DbQuery{
		index: "cabs",
		indexType: "cab",
		distance: "30km",
		etaModifier: 1.5,
		cabSize: 3,
	}
}

func (finder *DbQuery) SetIndex(index string) *DbQuery {
	finder.index = index
	return finder
}

func(finder *DbQuery) GetEta(lat, lon float64, vacant bool) float64 {
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("vacant", vacant))
	distanceQuery := elastic.NewGeoDistanceQuery("location").Distance(finder.distance).Point(lat, lon)
	query = query.Filter(distanceQuery)
	sorter := elastic.NewGeoDistanceSort("location").Point(lat, lon).Unit("km").GeoDistance("sloppy_arc").Asc()
	result, err := database.Search().Index(finder.index).Type(finder.indexType).
		Query(query).SortBy(sorter).Size(finder.cabSize).Do()
	HandleError(err)
	distance_sum := 0.0
	if result.Hits != nil {
		for _, hit := range result.Hits.Hits {
			distance_sum += hit.Sort[0].(float64)
		}
	}
	return distance_sum / float64(len(result.Hits.Hits)) * finder.etaModifier
}