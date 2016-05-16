package main

import (
	"gopkg.in/olivere/elastic.v3"
)

var database struct {
	source *elastic.Client
}

func InitDatabase() {
	client, err := elastic.NewClient()
	HandleError(err)
	database.source = client
}

func GetEta(lat, lon float64) float64 {
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("vacant", true))
	distance := elastic.NewGeoDistanceQuery("location").Distance("30km").Point(lat, lon)
	query = query.Filter(distance)
	sorter := elastic.NewGeoDistanceSort("location").Point(lat, lon).Unit("km").GeoDistance("sloppy_arc").Asc()
	result, err := database.source.Search().Index("cabs").Type("cab").Query(query).SortBy(sorter).Size(3).Pretty(true).Do()
	HandleError(err)
	distance_sum := 0.0
	if result.Hits != nil {
		for _, hit := range result.Hits.Hits {
			distance_sum += hit.Sort[0].(float64)
		}
	}
	return distance_sum / float64(len(result.Hits.Hits)) * 1.5
}
