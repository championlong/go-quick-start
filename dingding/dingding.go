package dingding

import (
	"github.com/championlong/backend-common/backend-common/dingding/model"
)

var dingdingConfig *model.DingdingConfig

func GetConfig() *model.DingdingConfig {
	return dingdingConfig
}

func Init(config model.DingdingConfig) {
	dingdingConfig = &config
}
