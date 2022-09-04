package factory

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type yamlRuleConfigParserFactory struct {
}

func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

type jsonRuleConfigParserFactory struct {
}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}