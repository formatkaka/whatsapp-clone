package v1

import (
	"sync"
	"whatsapp-clone/db"

	"github.com/olivere/elastic/v7"
	"go.elastic.co/apm/module/apmgoredis"
)

type SampleModule struct {
	controller *SampleController
	service    *SampleService
}

var sampleModuleSingleton *SampleModule
var sampleModuleOnce sync.Once

func NewSampleModule(dbFactory db.DBFactory, redis apmgoredis.Client, es *elastic.Client) *SampleModule {
	sampleModuleOnce.Do(func() {
		service := NewSampleService(dbFactory, redis, es)
		controller := NewSampleController(service)

		sampleModuleSingleton = &SampleModule{
			service:    service,
			controller: controller,
		}
	})

	return sampleModuleSingleton
}

func (s *SampleModule) GetController() *SampleController {
	return s.controller
}
