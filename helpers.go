package main

import (
	"fmt"
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
