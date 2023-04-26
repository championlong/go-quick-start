package system

import "github.com/championlong/go-quick-start/app/service"

type ApiGroup struct {
	BaseApi
	ChatGptApi
}

var (
	jwtService     = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService    = service.ServiceGroupApp.SystemServiceGroup.UserService
	chatGptService = service.ServiceGroupApp.SystemServiceGroup.ChatGptService
)
