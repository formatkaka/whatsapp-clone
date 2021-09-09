package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "pong"})
}
