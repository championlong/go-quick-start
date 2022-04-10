package v1

import (
	"github.com/championlong/backend-common/app/controller/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
