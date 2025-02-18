package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meingga/mrt-schedules/modules/stations"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()
	api := router.Group("/v1/api")

	stations.Initiate(api)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
