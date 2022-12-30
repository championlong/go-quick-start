package dingding

import (
	"github.com/championlong/backend-common/dingding/config"
	info "github.com/championlong/backend-common/dingding/message"
)

var dingdingConfig *config.DingdingConfig

func GetConfig() *config.DingdingConfig {
	return dingdingConfig
}

func Init(config config.DingdingConfig) {
	dingdingConfig = &config
}

func SendMessage(kindRobot string, messageType info.MsgType, message interface{}, at info.At) error {
	job := new(info.DingdingMasterJob)
	job.KindRobot = kindRobot
	job.Url = GetConfig().Url
	job.Msgtype = messageType
	job.Query = GetConfig().DingdingQuery[kindRobot]
	job.CommonMessage = message
	job.At = at
	return job.SendMessage()
}
