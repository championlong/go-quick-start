package alert

import (
	"github.com/championlong/backend-common/design_patterns/ocp/domain"
)

type ErrorAlertHandler struct{
	rule domain.AlertRule
	notification domain.Notification
}

func NewErrorAlertHandler(rule domain.AlertRule, notification domain.Notification) *ErrorAlertHandler {
	return &ErrorAlertHandler{
		rule:         rule,
		notification: notification,
	}
}

func (self *ErrorAlertHandler) Check(apiStatInfo ApiStatInfo) {
	if apiStatInfo.ErrorCount > self.rule.GetErrCount(apiStatInfo.Api) {
		self.notification.Notify("Error 超出阈值！")
	}
}
