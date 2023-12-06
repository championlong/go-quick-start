package app

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/championlong/go-quick-start/internal/app/global"
	"github.com/championlong/go-quick-start/internal/app/service/system"
	"github.com/championlong/go-quick-start/internal/pkg/initialize"
	"github.com/championlong/go-quick-start/pkg/log"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func Run() error {
	// root 适配性
	// 根据root位置去找到对应迁移位置,保证root路径有效
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)),
	)

	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	log.Info("server run success on ", zap.String("address", address))

	return s.ListenAndServe()
}
