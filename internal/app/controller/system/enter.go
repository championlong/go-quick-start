package system

import (
	"github.com/championlong/go-quick-start/internal/app/service"
)

type ApiGroup struct {
	BaseApi
	ChatGptApi
	DingRoles
}

var (
	jwtService     = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService    = service.ServiceGroupApp.SystemServiceGroup.UserService
	chatGptService = service.ServiceGroupApp.SystemServiceGroup.ChatGptService
)
