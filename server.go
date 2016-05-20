package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetFloat64(context *gin.Context, paramName string) (float64, error) {
	result, err := strconv.ParseFloat(context.Query(paramName), 64)
	return result, err
}
func EtaApi(c *gin.Context) {
	lat, err := GetFloat64(c, "lat")
	if err != nil {
		SendErrorBack(c, "lat is missing", 1)
		return
	}
	lon, err := GetFloat64(c, "lon")
	if err != nil {
		SendErrorBack(c, "lon is missing", 2)
		return
	}
	eta := NewDbQuery(IndexName).GetEta(lat, lon, true)

	c.JSON(200, gin.H{
		"eta": eta,
	})
}

func SendErrorBack(c *gin.Context, message string, code int) {
	c.JSON(500, gin.H{
		"error": gin.H{
			"message" : message,
			"code": code,
		},
	})
}

func StartEtaServer(){
	router := gin.Default()
	router.GET("/api/v1/cabs/eta", EtaApi)
	LogFatal(router.Run(":3000"))
}
