package v2

import "sync"

type SampleModule struct {
	controller *SampleController
	service    *SampleService
}

var sampleModuleSingleton *SampleModule
var sampleModuleOnce sync.Once

func NewSampleModule() *SampleModule {
	sampleModuleOnce.Do(func() {
		service := NewSampleService()
		controller := NewSampleController()
		sampleModuleSingleton = &SampleModule{
			service:    service,
			controller: controller,
		}
	})
	return sampleModuleSingleton
}

func (m *SampleModule) GetController() *SampleController {
	return m.controller
}
