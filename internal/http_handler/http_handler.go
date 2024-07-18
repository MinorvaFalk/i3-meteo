package http_handler

import (
	"i3/internal/model"
	"i3/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler interface {
	ScheduleLocationData(ctx *gin.Context)
	FetchWeathersData(c *gin.Context)
}

type httpHandler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) HTTPHandler {
	return &httpHandler{
		repo: repo,
	}
}

func (h *httpHandler) ScheduleLocationData(c *gin.Context) {
	var req model.WeatherRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, name, lat, lon, err := h.repo.InsertWeatherJob(c.Request.Context(), req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data: map[string]any{
			"id":   id,
			"name": name,
			"lat":  lat,
			"lon":  lon,
		},
	})
}

func (h *httpHandler) FetchWeathersData(c *gin.Context) {
	data, err := h.repo.FetchWeatherJobs(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    data,
	})
}
