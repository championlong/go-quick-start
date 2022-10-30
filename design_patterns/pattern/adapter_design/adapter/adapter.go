package adapter

import "strings"

type ASensitiveWordsFilter struct {
}

func (self ASensitiveWordsFilter) filterSexyWords(text string) string {
	return strings.Replace(text, "性感", "***", -1)
}

type BSensitiveWordsFilter struct {
}

func (self BSensitiveWordsFilter) filterPoliticalWords(text string) string {
	return strings.Replace(text, "政治", "***", -1)
}

type ISensitiveWordsFilter interface {
	filter(text string) string
}

type ASensitiveWordsFilterAdaptor struct {
	ASensitiveWordsFilter
}

func (self ASensitiveWordsFilterAdaptor) filter(text string) string {
	return self.filterSexyWords(text)
}

type BSensitiveWordsFilterAdaptor struct {
	BSensitiveWordsFilter
}

func (self BSensitiveWordsFilterAdaptor) filter(text string) string {
	return self.filterPoliticalWords(text)
}

type RiskManagement struct {
	filters []ISensitiveWordsFilter
}

func (self *RiskManagement) addSensitiveWordsFilter(filter ISensitiveWordsFilter) {
	self.filters = append(self.filters, filter)
}

func (self *RiskManagement) filterSensitiveWords(text string) string {
	for _, filter := range self.filters {
		text = filter.filter(text)
	}
	return text
}
