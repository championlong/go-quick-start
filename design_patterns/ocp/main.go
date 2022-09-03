package main

import (
	"github.com/championlong/backend-common/design_patterns/ocp/alert"
	"github.com/championlong/backend-common/design_patterns/ocp/initialize"
)

/*
开闭原则
添加一个新的功能应该是，在已有代码基础上扩展代码(新增模块、类、方法等)，而非修改已有代码(修改模块、类、方法等)
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
