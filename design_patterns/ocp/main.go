package main

import (
	"github.com/championlong/backend-common/design_patterns/ocp/alert"
	"github.com/championlong/backend-common/design_patterns/ocp/initialize"
)

/*
开闭原则:
对扩展开放、修改关闭
添加一个新的功能应该是，在已有代码基础上扩展代码(新增模块、类、方法等)，而非修改已有代码(修改模块、类、方法等)
最常用来提高代码扩展性的方法有:多态、依赖注入、基于接口而非实现编程，以及大部分的设计模式(比如，装饰、策略、模 板、职责链、状态)。
 */
func main() {
	apiStatInfo := alert.ApiStatInfo{
		Api:               "test",
		RequestCount:      50,
		ErrorCount:        20,
		DurationOfSeconds: 2,
	}
	initialize.GetApplicationContext().GetAlert().Check(apiStatInfo)
}
