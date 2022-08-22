package initialize

import (
	"github.com/championlong/backend-common/design_patterns/ocp/alert"
	"github.com/championlong/backend-common/design_patterns/ocp/domain"
)

type ApplicationContext struct {
	alertRule    domain.AlertRule
	notification domain.Notification
	Alert        alert.Alert
}

func (self *ApplicationContext) initializeBeans() {
	self.alertRule = domain.AlertRule{}
	self.notification = domain. Notification{}
	self.Alert.AddAlertHandler(alert.NewTpsAlertHandler(self.alertRule, self.notification))
	self.Alert.AddAlertHandler(alert.NewErrorAlertHandler(self.alertRule, self.notification))
}

func (self *ApplicationContext) GetAlert() *alert.Alert {
	return &self.Alert
}

var Instance *ApplicationContext

func GetApplicationContext() *ApplicationContext {
	if Instance == nil {
		Instance = new(ApplicationContext)
		Instance.initializeBeans()
	}
	return Instance
}
