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
	client *elastic.Client
}

type Cab struct {

}

func NewDbQuery(index string) *DbQuery {
	return &DbQuery{
		index: index,
		indexType: "cab",
		distance: "30km",
		etaModifier: 1.5,
		cabSize: 3,
		client: database,
	}
}

func (finder *DbQuery) PutMappings() *DbQuery {
	mapping := `{
		"cab":{
			"properties":{
				"location":{
					"type":"geo_point"
				},
				"vacant":{
					"type":"boolean"
				}
			}
		}
	}`
	_, err := finder.client.PutMapping().Index(finder.index).Type(finder.indexType).BodyString(mapping).Do()
	HandleError(err)
	return finder
}

func (finder *DbQuery) CreateIndex() *DbQuery {
	_, err := finder.client.CreateIndex(finder.index).Do()
	HandleError(err)
	return finder
}

func (finder *DbQuery) DestroyIndex() *DbQuery {
	_, err := finder.client.DeleteIndex(finder.index).Do()
	HandleError(err)
	return finder
}

func (finder *DbQuery) IndexExists() bool {
	exist, err := finder.client.IndexExists(finder.index).Do()
	if err != nil {
		return false
	} else {
		return exist
	}
}


func(finder *DbQuery) GetEta(lat, lon float64, vacant bool) float64 {
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
	return distance_sum / float64(len(result.Hits.Hits)) * finder.etaModifier
}