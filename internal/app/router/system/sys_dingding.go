package system

import (
	v1 "github.com/championlong/go-quick-start/internal/app/controller"
	"github.com/championlong/go-quick-start/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type DingTalkRouter struct{}

func (s *ChatGptRouter) InitDingTalkRouter(Router *gin.RouterGroup) {
	chatGptRouter := Router.Group("ding").Use(middleware.OperationRecord())
	chatGptApi := v1.ApiGroupApp.SystemApiGroup.DingRoles
	{
		chatGptRouter.POST("roles", chatGptApi.DingApplyRoles)
	}
}
