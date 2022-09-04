package dao

import (
	"fmt"
	"github.com/championlong/backend-common/design_patterns/counter/vo"
)

var RequestInfos map[string][]vo.RequestInfo

// MetricsStorage 面向接口编程，后续存储方式只需实现接口即可，其他接口函数调用的地方都不需要改动，满足开闭原则。
type MetricsStorage interface {
	SaveRequestInfo(requestInfo vo.RequestInfo)
	GetRequestInfosByDuration(apiName string, startTimestamp int64, endTimestamp int64) []vo.RequestInfo
	GetAllRequestInfosByDuration(startTimestamp, endTimestamp float64) map[string][]vo.RequestInfo
}

type RedisMetricsStorage struct {
}

func (self *RedisMetricsStorage) SaveRequestInfo(requestInfo vo.RequestInfo) {
	RequestInfos[requestInfo.ApiName] = append(RequestInfos[requestInfo.ApiName], requestInfo)
	fmt.Println("SaveRequestInfo")
}

func (self *RedisMetricsStorage) GetRequestInfosByDuration(apiName string, startTimestamp int64, endTimestamp int64) []vo.RequestInfo {
	fmt.Println("GetRequestInfosByDuration")
	return []vo.RequestInfo{}
}

func (self *RedisMetricsStorage) GetAllRequestInfosByDuration(startTimestamp, endTimestamp float64) map[string][]vo.RequestInfo {
	fmt.Println("GetAllRequestInfosByDuration")
	return RequestInfos
}
