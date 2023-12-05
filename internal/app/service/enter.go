package service

import (
	"github.com/championlong/go-quick-start/internal/app/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
