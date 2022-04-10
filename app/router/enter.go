package router

import "github.com/championlong/backend-common/app/router/system"

type RouterGroup struct {
	System   system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
