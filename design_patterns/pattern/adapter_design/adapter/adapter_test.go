package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	var filter RiskManagement
	filter.addSensitiveWordsFilter(new(ASensitiveWordsFilterAdaptor))
	filter.addSensitiveWordsFilter(new(BSensitiveWordsFilterAdaptor))
	text := "其中包含的屏蔽字符：性感，政治"
	fmt.Println(filter.filterSensitiveWords(text))
}