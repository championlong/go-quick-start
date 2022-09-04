package service

import "github.com/championlong/backend-common/design_patterns/counter/dao"

type EmailReporter struct {
	ScheduledReporter
}

func NewEmailReporter(aggregator Aggregator, metricsStorage dao.MetricsStorage, viewer StatViewer) *EmailReporter {
	return &EmailReporter{
		ScheduledReporter{
			metricsStorage: metricsStorage,
			aggregator:     aggregator,
			viewer:         viewer,
		},
	}
}

func (self *EmailReporter) startDailyReport() {

}
