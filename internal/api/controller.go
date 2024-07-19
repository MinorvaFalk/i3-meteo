package api

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	ScheduleLocationData(ctx *gin.Context)

	FetchWeathersData(c *gin.Context)
	FetchCities(c *gin.Context)
}
