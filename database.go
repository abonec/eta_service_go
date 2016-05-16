package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v3"
)

type Database struct {
	source *elastic.Client
}

func (v *Database) Init() {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	v.source = client
}

func PrintQuery(query elastic.Query) {
	src, err := query.Source()
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
func (v *Database) GetEta() float64 {
	lat := 55.662987
	lon := 37.656230
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("vacant", true))
	distance := elastic.NewGeoDistanceQuery("location").Distance("30km").Point(lat, lon)
	query = query.Filter(distance)
	sorter := elastic.NewGeoDistanceSort("location").Point(lat, lon).Unit("km").GeoDistance("sloppy_arc").Asc()
	result, err := v.source.Search().Index("cabs").Type("cab").Query(query).SortBy(sorter).Size(3).Pretty(true).Do()
	if err != nil {
		panic(err)
	}
	distance_sum := 0.0
	if result.Hits != nil {
		for _, hit := range result.Hits.Hits {
			distance_sum += hit.Sort[0].(float64)
		}
	}
	return distance_sum / float64(len(result.Hits.Hits))
}

func InitDatabase() *Database {
	database := &Database{}
	database.Init()
	return database
}
