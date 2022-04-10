package service

import (
	"github.com/championlong/backend-common/app/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
