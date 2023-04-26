package system

import "github.com/championlong/go-quick-start/app/service"

type ApiGroup struct {
	JwtApi
	BaseApi
	DBApi
}

var (
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
