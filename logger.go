package main

import "log"

func LogInfo(format string, v ...interface{}) {
	log.Printf(format, v...)
}
