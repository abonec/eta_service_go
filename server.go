package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Eta(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	HandleError(err)
	lon, err := strconv.ParseFloat(c.Query("lon"), 64)
	HandleError(err)
	eta := NewDbQuery().GetEta(lat, lon, true)

	c.JSON(200, gin.H{
		"eta": eta,
	})
}

func main() {
	router := gin.Default()
	router.GET("/api/v1/cabs/eta", Eta)
	InitDatabase()
	log.Fatal(router.Run(":3000"))
}
