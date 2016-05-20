package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func EtaApi(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	HandleError(err)
	lon, err := strconv.ParseFloat(c.Query("lon"), 64)
	HandleError(err)
	eta := NewDbQuery(IndexName).GetEta(lat, lon, true)

	c.JSON(200, gin.H{
		"eta": eta,
	})
}

func StartEtaServer(){
	router := gin.Default()
	router.GET("/api/v1/cabs/eta", EtaApi)
	LogFatal(router.Run(":3000"))
}
