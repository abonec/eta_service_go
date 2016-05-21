package main

import (
	"flag"
	"gopkg.in/olivere/elastic.v3"
)

const (
	IndexName = "cabs"
	cab_fixtures_file_path = "fixtures/cabs.txt"
	cab_fixtures_size = 1000
	messager_cab_queue_name = "cab_queue"
	messager_cab_updated_name = "cab_updated"
	messager_cab_updated_exchange_name = "cab_updated_exchange"
	messager_queue_url = "amqp://guest:guest@localhost:5672/"
)

var (
	mode = flag.String("mode", "eta_server", "start in mode")
	message = flag.String("message", "hello world", "message to send via message queue")
	database *elastic.Client
	amqpMessager *AMQPMessager
	amqpSender *AMQPSender
	amqpCabUpdatedExchange *AMQPFanoutExchange
)