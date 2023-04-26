package router

import "github.com/championlong/go-quick-start/app/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
