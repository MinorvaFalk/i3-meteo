package apicontroller

import (
	"i3/internal/api"
	"i3/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	uc api.Usecase
}

func New(uc api.Usecase) api.Controller {
	return &controller{
		uc: uc,
	}
}

func (h *controller) ScheduleLocationData(c *gin.Context) {
	var req model.CityRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.uc.ScheduleLocationData(c.Request.Context(), req)
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

func (h *controller) FetchWeathersData(c *gin.Context) {
	data, err := h.uc.FetchWeatherData(c.Request.Context())
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
