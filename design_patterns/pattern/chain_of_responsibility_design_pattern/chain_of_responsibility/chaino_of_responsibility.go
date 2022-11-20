package chain_of_responsibility

type SensitiveWordFilter interface {
	doFilter(content string) bool
}

// SensitiveWordFilterChain 职责链
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

// AddFilter 添加一个过滤器
func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

// Filter 执行过滤
func (c *SensitiveWordFilterChain) doFilter(content string) bool {
	for _, filter := range c.filters {
		if filter.doFilter(content) {
			return true
		}
	}
	return false
}

type AdSensitiveWordFilter struct{}

func (f *AdSensitiveWordFilter) doFilter(content string) bool {
	return false
}

type PoliticalWordFilter struct{}

func (f *PoliticalWordFilter) doFilter(content string) bool {
	return true
}
