package service

import "github.com/championlong/backend-common/design_patterns/counter/vo"

type Aggregator struct {
}

func (self *Aggregator) Aggregate(requestInfos map[string][]vo.RequestInfo, durationInMillis float64) map[string]vo.RequestStat {
	requestStats := make(map[string]vo.RequestStat)
	for apiName, requestInfosPerApi := range requestInfos {
		requestStat := self.doAggregate(requestInfosPerApi, durationInMillis)
		requestStats[apiName] = requestStat
	}
	return requestStats
}

func (self *Aggregator) doAggregate(requestInfos []vo.RequestInfo, durationInMillis float64) vo.RequestStat {
	var respTimes []float64
	for _, requestInfo := range requestInfos {
		respTime := requestInfo.ResponseTime
		respTimes = append(respTimes, respTime)
	}
	var requestStat vo.RequestStat
	requestStat.MaxResponseTime = self.max(respTimes)
	requestStat.MinResponseTime = self.min(respTimes)
	requestStat.AvgResponseTime = self.avg(respTimes)
	requestStat.P999ResponseTime = self.percentile999(respTimes)
	requestStat.P99ResponseTime = self.percentile99(respTimes)
	requestStat.Tps = self.tps(len(respTimes), durationInMillis/1000)
	return requestStat
}

func (self *Aggregator) max(dataset []float64) float64 {
	var max float64
	for _, v := range dataset {
		if v > max {
			max = v
		}
	}
	return max
}
func (self *Aggregator) min(dataset []float64) float64 {
	var min float64
	for _, v := range dataset {
		if v < min {
			min = v
		}
	}
	return min
}
func (self *Aggregator) avg(dataset []float64) float64 {
	if len(dataset) == 0 {
		return 0
	}
	var count float64
	for _, v := range dataset {
		count += v
	}
	return count / float64(len(dataset))
}
func (self *Aggregator) tps(count int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(count) / duration
}
func (self *Aggregator) percentile999(dataset []float64) float64 {
	return 0
}
func (self *Aggregator) percentile99(dataset []float64) float64 {
	return 0
}
func (self *Aggregator) percentile(dataset []float64, ratio float64) float64 {
	return 0
}
