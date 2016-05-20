package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

type paramMethod func(string) string
func GetFloat64(context *gin.Context, paramName string, method paramMethod) (float64, error) {
	result, err := strconv.ParseFloat(method(paramName), 64)
	return result, err
}
func GetBool(context *gin.Context, paramName string, method paramMethod) (bool, error) {
	result, err := strconv.ParseBool(method(paramName))
	return result, err
}
func EtaApi(c *gin.Context) {
	lat, err := GetFloat64(c, "lat", c.Query)
	if err != nil {
		SendErrorBack(c, "lat is missing", 1)
		return
	}
	lon, err := GetFloat64(c, "lon", c.Query)
	if err != nil {
		SendErrorBack(c, "lon is missing", 1)
		return
	}
	eta := NewDbQuery(IndexName).GetEta(lat, lon, true)

	c.JSON(200, gin.H{
		"eta": eta,
	})
}

func CabsUpdate(c *gin.Context) {
	vacant, err := GetBool(c, "vacant", c.PostForm)
	if MissingParam(c, err, "vacant", 1) {
		return
	}
	lat, err := GetFloat64(c, "lat", c.PostForm)
	if MissingParam(c, err, "lat", 1) {
		return
	}
	lon, err := GetFloat64(c, "lon", c.PostForm)
	if MissingParam(c, err, "lon", 1) {
		return
	}

	cab := &Cab{
		Vacant: vacant,
		Location: Location{
			Lat: lat,
			Lon: lon,
		},

	}
	amqpSender.Send(cab.ToJson())

	c.JSON(200, gin.H{
		"result": "ok",
		"queue": "sent",
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

func MissingParam(context *gin.Context, err error, paramName string, code int) bool {
	if err != nil {
		error_message := fmt.Sprintf("%s is missing", paramName)
		SendErrorBack(context, error_message, code)
		return true
	} else {
		return false
	}
}

func StartEtaServer(){
	router := gin.Default()
	router.GET("/api/v1/cabs/eta", EtaApi)
	LogFatal(router.Run(":3000"))
}

func StartUpdateCabServer(){
	router := gin.Default()
	router.PUT("/api/v1/cabs", CabsUpdate)
	LogFatal(router.Run(":3001"))
}