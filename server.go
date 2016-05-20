package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"flag"
)

func Eta(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	HandleError(err)
	lon, err := strconv.ParseFloat(c.Query("lon"), 64)
	HandleError(err)
	eta := NewDbQuery(IndexName).GetEta(lat, lon, true)

	c.JSON(200, gin.H{
		"eta": eta,
	})
}

func init() {
	flag.Parse()
	InitDatabase()
	InitLogger()
	InitMessageQueue()
}
func main() {
	switch *mode {
	case "server":
		router := gin.Default()
		router.GET("/api/v1/cabs/eta", Eta)
		LogFatal(router.Run(":3000"))
	case "migrate":
		NewDbQuery(IndexName).Migrate(-1)
	case "send_message":
		SendMessage([]byte(*message))
	}
}
