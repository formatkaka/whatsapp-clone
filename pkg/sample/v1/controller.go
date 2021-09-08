package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleController struct {
	service *SampleService
}

func NewSampleController(service *SampleService) *SampleController {
	return &SampleController{
		service: service,
	}
}

func (s *SampleController) Hello(c *gin.Context) {
	resp := s.service.Hello()
	c.JSON(http.StatusOK, gin.H{resp: resp})
}

func (s *SampleController) Error(c *gin.Context) {
	err := errors.New("error message")
	c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
}

func (s *SampleController) FromDb(c *gin.Context) {
	users, err := s.service.FromDb(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (s *SampleController) FromRedis(c *gin.Context) {
	info, err := s.service.FromRedis(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"info": info})
}

func (s *SampleController) Panic(c *gin.Context) {
	panic("ggwp")
}

func (s *SampleController) FromEs(c *gin.Context) {
	info, err := s.service.FromES()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"result": info})
}
