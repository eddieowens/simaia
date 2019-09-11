package app

import (
	"github.com/eddieowens/axon"
	"github.com/eddieowens/simaia/app/service"
)

func CreateInjector() axon.Injector {
	return axon.NewInjector(axon.NewBinder(
		new(service.Package),
	))
}
