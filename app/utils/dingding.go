package utils

import (
	"github.com/championlong/backend-common/app/constants"
	"github.com/championlong/backend-common/dingding"
)

func SendDingdingAlertError(content string) {
	err := dingding.DingDingSend.GetConfigWebHookRobot(constants.ALERT_WEBHOOK_ROBOT).DingDingSendGroupText(content, nil, nil, false)
	if err != nil {
		Err("DingDingSend alert_robot is err:%s", err.Error())
	}
}
