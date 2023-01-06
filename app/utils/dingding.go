package utils

import (
	"github.com/championlong/backend-common/app/constants"
	dingding "github.com/championlong/dingtalk-sdk"
	"github.com/championlong/dingtalk-sdk/model"
)

func SendDingAlertError(content string) {
	textMessage := model.TextMessage{Content: content}
	err := dingding.SendDingMessage(constants.ALERT_WEBHOOK_ROBOT, dingding.MsgTypeText, textMessage, model.At{})
	if err != nil {
		Err("DingDingSend alert_robot is err:%s", err.Error())
	}
}
