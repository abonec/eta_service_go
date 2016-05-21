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
	case "update_cab_server":
		StartUpdateCabServer()
	case "update_server_log":
		StartUpdateServerLog()
	case "migrate":
		Migrate()
	case "send_message":
		SendMessage([]byte(*message))
	}
}
