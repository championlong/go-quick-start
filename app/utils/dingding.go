package utils

import (
	"github.com/championlong/go-quick-start/app/constants"
	"github.com/championlong/go-quick-start/dingding"
)

func SendDingdingAlertError(content string) {
	err := dingding.DingDingSend.GetConfigWebHookRobot(constants.ALERT_WEBHOOK_ROBOT).DingDingSendGroupText(content, nil, nil, false)
	if err != nil {
		Err("DingDingSend alert_robot is err:%s", err.Error())
	}
}
