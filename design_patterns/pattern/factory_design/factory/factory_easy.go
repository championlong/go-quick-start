package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (J jsonRuleConfigParser) Parse(data []byte) {

}

type yamlRuleConfigParser struct {
}

func (Y yamlRuleConfigParser) Parse(data []byte) {

}

// 简单工厂模式
func NewIRuleConfigParser(configFormat string) IRuleConfigParser {
	switch configFormat {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}

// 单例简单工厂模式，要提前初始化map，将实例放入map中
var cachedParsers map[string]IRuleConfigParser
func NewIRuleConfigParser2(configFormat string) IRuleConfigParser {
	return cachedParsers[configFormat]
}
