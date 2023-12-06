package middleware

import (
	"github.com/championlong/go-quick-start/internal/app/global"
	"github.com/championlong/go-quick-start/internal/app/model/common/response"
	"github.com/championlong/go-quick-start/internal/app/service"
	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup

// 拦截器.
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		// obj := c.Request.URL.Path
		// 获取请求方法
		// act := c.Request.Method
		// 获取用户的角色
		// sub := waitUse.AuthorityId
		// e := casbinService.Casbin()
		// 判断策略中是否存在
		// success, _ := e.Enforce(sub, obj, act)
		if global.GVA_CONFIG.System.Env == "develop" {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
