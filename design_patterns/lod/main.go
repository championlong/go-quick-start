package main

import (
	"fmt"
	"github.com/championlong/backend-common/design_patterns/lod/download"
)

/*
迪米特法则(LOD): 最小知识原则
高内聚: 用来指导类本身的设计，是指相近的功能应该放到同一个类中，不相近的功能不要放到同一个类中。相近的功能往往 会被同时修改，放到同一个类中，修改会比较集中，代码容易维护。
松耦合: 用来指导类与类之间依赖关系的设计，在代码中，类与类之间的依赖关系简单清晰。即使两个类有依赖关系，一个类的代码改动不会或者很少导致依赖类的代码改动。
解耦:
1.封装与抽象
2.中间层
3.模块化
4.设计思想和原则(单一职责原则、基于接口而非实现编程、依赖注入、多用组合少用继承、迪米特法则)
每个模块(unit)只应该了解那些与它关系密切的模块(units: only units “closely” related to the current unit)的有限知识(knowledge)。或者说，每个模块只和自己的朋 友“说话”(talk)，不和陌生人“说话”(talk)
不该有直接依赖关系的类之间，不要有依赖;有依赖关系的类之间，尽量只依赖必要的接口
*/
func main() {
	//不该有直接依赖关系的类之间，不要有依赖
	htmlDownloader := download.HtmlDownloader{}
	document, err := download.NewDocumentFactory(htmlDownloader).CreateDocument("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(document)
}
