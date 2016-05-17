package main

import (
	"gopkg.in/olivere/elastic.v3"
)

var database *elastic.Client

func InitDatabase() {
	client, err := elastic.NewClient()
	HandleError(err)
	database = client
}

