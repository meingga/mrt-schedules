package stations

import (
	"github.com/gin-gonic/gin"
	"github.com/meingga/mrt-schedules/common/response"
	"net/http"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	stations := router.Group("/stations")
	stations.GET("", func(c *gin.Context) {
		//	code service
		GetAllStations(c, stationService)
	})

	stations.GET("/:id", func(c *gin.Context) {
		CheckSchedulesByStation(c, stationService)
	})
}

func GetAllStations(c *gin.Context, service Service) {
	data, err := service.GetAllStations()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully fetched stations",
		Data:    data,
	})
}

func CheckSchedulesByStation(c *gin.Context, service Service) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "ID is required",
			Data:    nil,
		})
		return
	}

	data, err := service.CheckScheduleByStation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully fetched schedules by station",
		Data:    data,
	})
}
