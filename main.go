package main

import (
	"flag"
)

func init() {
	flag.Parse()
	InitDatabase()
	InitLogger()
	InitMessageQueue()
}
func main() {
	switch *mode {
	case "eta_server":
		StartEtaServer()
	case "migrate":
		Migrate()
	case "send_message":
		SendMessage([]byte(*message))
	}
}
