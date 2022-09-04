package vo

type RequestInfo struct {
	ApiName      string
	ResponseTime float64
	Timestamp    int64
}

func NewRequestInfo(apiName string, responseTime float64, timestamp int64) RequestInfo {
	return RequestInfo{
		ApiName:      apiName,
		ResponseTime: responseTime,
		Timestamp:    timestamp,
	}
}

type RequestStat struct {
	MaxResponseTime  float64
	MinResponseTime  float64
	AvgResponseTime  float64
	P999ResponseTime float64
	P99ResponseTime  float64
	Count            float64
	Tps              float64
}
