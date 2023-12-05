package v1

import (
	"github.com/championlong/go-quick-start/internal/app/controller/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
