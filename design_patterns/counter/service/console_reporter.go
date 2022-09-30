package service

import (
	"fmt"
	"github.com/championlong/backend-common/design_patterns/counter/dao"
	"time"
)

// ConsoleReporter 把核心逻辑都剥离出去，形成独立的类，上帝类只负责组装类和串联执行流程。这样做的好处是，代码结构更加清晰，底层核心逻辑更容易被复用。
type ConsoleReporter struct {
	ScheduledReporter
}

func NewConsoleReporter(metricsStorage dao.MetricsStorage, aggregator Aggregator, viewer StatViewer) *ConsoleReporter {
	return &ConsoleReporter{
		ScheduledReporter{
			metricsStorage: metricsStorage,
			aggregator:     aggregator,
			viewer:         viewer,
		},
	}
}

func NewConsoleRedisReporter() *ConsoleReporter {
	return &ConsoleReporter{
		ScheduledReporter{
			metricsStorage: &dao.RedisMetricsStorage{},
			aggregator:     Aggregator{},
			viewer:         &ConsoleViewer{},
		},
	}
}

func (self *ConsoleReporter) StartRepeatedReport(periodInSeconds time.Duration, durationInSeconds float64) {
	ticker := time.NewTicker(periodInSeconds * time.Second)
	for _ = range ticker.C {
		fmt.Println("start")
		durationInMillis := durationInSeconds * 1000
		endTimeInMillis := float64(time.Now().UnixMilli())
		startTimeInMillis := endTimeInMillis - durationInMillis
		self.doStatAndReport(startTimeInMillis, endTimeInMillis)
	}
}
