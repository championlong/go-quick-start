package collector

import (
	"github.com/championlong/backend-common/design_patterns/counter/dao"
	"github.com/championlong/backend-common/design_patterns/counter/vo"
)

// MetricsCollector 面向接口而非实现
type MetricsCollector struct {
	metricsStorage dao.MetricsStorage
}

// NewMetricsCollector 依赖注入的方式来传递MetricsStorage对象，后续可以灵活地替换不同的存储方式，满足开闭原则。
func NewMetricsCollector(metricsStorage dao.MetricsStorage) *MetricsCollector {
	return &MetricsCollector{
		metricsStorage: metricsStorage,
	}
}

// NewMetricsRedisCollector 框架的易用性
func NewMetricsRedisCollector() *MetricsCollector {
	return &MetricsCollector{
		metricsStorage: &dao.RedisMetricsStorage{},
	}
}

func (self *MetricsCollector) RecordRequest(requestInfo vo.RequestInfo) {
	if requestInfo.ApiName == "" {
		return
	}
	self.metricsStorage.SaveRequestInfo(requestInfo)
}
