package alert

type AlertHandler interface {
	Check(apiStatInfo ApiStatInfo)
}

type ApiStatInfo struct {
	Api               string
	RequestCount      int64
	ErrorCount        int64
	DurationOfSeconds int64
}

type Alert struct {
	AlertHandlers []AlertHandler
}

func (self *Alert) AddAlertHandler(alertHandler AlertHandler) {
	self.AlertHandlers = append(self.AlertHandlers, alertHandler)
}

func (self *Alert) Check(apiStatInfo ApiStatInfo) {
	for _, v := range self.AlertHandlers {
		v.Check(apiStatInfo)
	}
}
