package router

import (
	"i3/config"
	"i3/internal/api"
	"i3/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	ginzap "github.com/gin-contrib/zap"
)

func NewGin(controller api.Controller) *gin.Engine {
	if config.ReadConfig().Env == "production" {
		gin.SetMode("release")
	}

	r := gin.Default()
	r.Use(ginzap.Ginzap(logger.Zap(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger.Zap(), true))
	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		c.Status(http.StatusInternalServerError)
	}))

	r.GET("/status", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	v1 := r.Group("/api/v1")

	v1.GET("/weathers", controller.FetchWeathersData)

	schedule := v1.Group("/schedule")
	schedule.POST("/location", controller.ScheduleLocationData)

	return r
}
