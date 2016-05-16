package main

import (
	"github.com/gin-gonic/gin"
	// "gopkg.in/olivere/elastic.v3"
	"log"
	"fmt"
)

func Eta(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")
	c.JSON(200, gin.H{
		"lat": lat,
		"lon": lon,
	})
}

func main() {
	router := gin.Default()
	router.GET("/api/v1/cabs/eta", Eta)

	fmt.Printf("%d", InitDatabase().GetEta())

	log.Fatal(router.Run(":3000"))
}
