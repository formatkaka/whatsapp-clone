package v1

import "sync"

type Module struct {
	controller *Controller
}

var moduleSingleton *Module
var moduleSingletonOnce sync.Once

func NewModule() *Module {
	moduleSingletonOnce.Do(func() {
		controller := NewController()

		moduleSingleton = &Module{
			controller: controller,
		}
	})

	return moduleSingleton
}

func (m *Module) GetController() *Controller {
	return m.controller
}
