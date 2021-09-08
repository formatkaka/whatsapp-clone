package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleController struct {
	service *SampleService
}

func NewSampleController() *SampleController {
	return &SampleController{}
}

func (s *SampleController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": s.service.Hello})
}
