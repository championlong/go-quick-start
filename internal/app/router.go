package app

import (
	"net/http"

	"github.com/championlong/go-quick-start/api/swagger/docs"
	"github.com/championlong/go-quick-start/internal/app/global"
	"github.com/championlong/go-quick-start/internal/app/router"
	"github.com/championlong/go-quick-start/internal/pkg/initialize"
	"github.com/championlong/go-quick-start/internal/pkg/middleware"
	"github.com/championlong/go-quick-start/pkg/log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go
	// 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	log.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	log.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup)     // 注册用户路由
		systemRouter.InitChatGptRouter(PrivateGroup)  // chatGpt接口
		systemRouter.InitDingTalkRouter(PrivateGroup) // 钉钉相关接口
	}

	initialize.InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	log.Info("router register success")
	return Router
}
