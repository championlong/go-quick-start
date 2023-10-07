package main

import (
	"github.com/championlong/go-quick-start/app/core"
	"github.com/championlong/go-quick-start/app/global"
	"github.com/championlong/go-quick-start/app/initialize"
	ding "github.com/championlong/robot-talk-sdk"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

//	@title						Swagger Example API
//	@version					0.0.1
//	@description				This is a sample Server pets
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						x-token
//	@BasePath					/
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	ding.Init(global.GVA_CONFIG.Dingding)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
