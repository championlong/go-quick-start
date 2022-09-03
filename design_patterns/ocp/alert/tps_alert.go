package alert

import (
	"github.com/championlong/backend-common/design_patterns/ocp/domain"
)

type TpsAlertHandler struct {
	rule         domain.AlertRule
	notification domain.Notification
}

func NewTpsAlertHandler(rule domain.AlertRule, notification domain.Notification) *TpsAlertHandler {
	return &TpsAlertHandler{
		rule:         rule,
		notification: notification,
	}
}
func (self *TpsAlertHandler) Check(apiStatInfo ApiStatInfo) {
	tps := apiStatInfo.RequestCount / apiStatInfo.DurationOfSeconds
	if tps > self.rule.GetTpsCount(apiStatInfo.Api) {
		self.notification.Notify("TPS 超出阈值！")
	}
}
