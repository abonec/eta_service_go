package main

import (
	"gopkg.in/olivere/elastic.v3"
)

func InitDatabase() {
	client, err := elastic.NewClient()
	HandleError(err)
	database = client
}
