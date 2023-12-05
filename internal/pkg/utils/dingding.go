package utils

import (
	"github.com/championlong/go-quick-start/internal/pkg/constants"
	"github.com/championlong/go-quick-start/internal/pkg/global"
	ding "github.com/championlong/robot-talk-sdk"
	"github.com/championlong/robot-talk-sdk/model/ding_talk"
	"github.com/championlong/robot-talk-sdk/platform"
	"go.uber.org/zap"
)

func SendDingdingAlertError(content string) {
	err := ding.SendDingMessage(constants.ALERT_WEBHOOK_ROBOT, platform.MsgTypeText, ding_talk.TextMessage{Content: content}, ding_talk.At{})
	if err != nil {
		global.GVA_LOG.Error("DingDingSend alert_robot is err!", zap.Error(err))
	}
}
