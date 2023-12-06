package main

import (
	"github.com/championlong/go-quick-start/internal/app"
	"github.com/championlong/go-quick-start/internal/app/global"
	"github.com/championlong/go-quick-start/internal/pkg/core"
	"github.com/championlong/go-quick-start/internal/pkg/initialize"
	"github.com/championlong/go-quick-start/pkg/log"
	ding "github.com/championlong/robot-talk-sdk"
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
	global.GVA_VP = core.Viper("configs/app_config.yaml", &global.GVA_CONFIG) // 初始化Viper
	log.Init(&global.GVA_CONFIG.Zap)
	defer log.Flush()
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
	app.NewApp("api").Run()
}
