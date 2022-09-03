package domain

type AlertRule struct {
}

func (self *AlertRule) GetTpsCount(api string) int64 {
	return 10
}

func (self *AlertRule) GetErrCount(api string) int64 {
	return 10
}