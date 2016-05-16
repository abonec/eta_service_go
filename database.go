package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v3"
)

type Database struct {
	source *elastic.Client
}

func (v Database) Init() {
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
func (v Database) GetEta() {
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("vacant", true))
	distance := elastic.NewGeoDistanceQuery("location").Distance("40 km").Point(55.662987, 37.656230)
	query = query.Filter(distance)
	PrintQuery(query)
	_, err := v.source.Search().Index("cabs").Type("cab").Query(query).Do()
	if err != nil {
		panic(err)
	}
}

func InitDatabase() Database {
	database := Database{}
	database.Init()
	return database
}
