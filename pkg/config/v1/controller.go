package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) APMSampleRate(c *gin.Context) {
	var body struct{ rate string }
	c.BindJSON(&body)
	controller.service.SetAPMSampleRate(body.rate)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
