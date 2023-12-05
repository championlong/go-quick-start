package main

import (
	core2 "github.com/championlong/go-quick-start/internal/pkg/core"
	"github.com/championlong/go-quick-start/internal/pkg/global"
	initialize2 "github.com/championlong/go-quick-start/internal/pkg/initialize"
	ding "github.com/championlong/robot-talk-sdk"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core2.Viper() // 初始化Viper
	global.GVA_LOG = core2.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	ding.Init(global.GVA_CONFIG.Dingding)
	global.GVA_DB = initialize2.Gorm() // gorm连接数据库
	initialize2.Timer()
	initialize2.DBList()
	if global.GVA_DB != nil {
		initialize2.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core2.RunWindowsServer()
}
