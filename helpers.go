package main

import (
	"fmt"
)

const (
	cab_fixtures_file_path = "fixtures/cabs.txt"
	cab_fixtures_size = 1000
)
func HandleError(error interface{}) {
	if error != nil {
		panic(error)
	}
}
func failOnError(err error, msg string) {
	if err != nil {
		LogFatal("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
