package v1

import (
	"sync"
)

type Module struct{
	service *Service
	controller *Controller
}

var moduleSingleton *Module
var moduleSingletonOnce sync.Once

func NewModuleSingleton() *Module {
	moduleSingletonOnce.Do(func() {
		service := NewService()
		controller := NewController(service)
		moduleSingleton = &Module{
			service: service,
			controller: controller,
		}
	})
	return moduleSingleton
}

func (m *Module) GetController() *Controller {
	return m.controller
}