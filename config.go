package main

import (
	"flag"
	"gopkg.in/olivere/elastic.v3"
)

const (
	IndexName = "cabs"
	cab_fixtures_file_path = "fixtures/cabs.txt"
	cab_fixtures_size = 1000
	cab_queue_name = "cab_queue"
)

var (
	mode = flag.String("mode", "eta_server", "start in mode")
	message = flag.String("message", "hello world", "message to send via message queue")
	database *elastic.Client
)