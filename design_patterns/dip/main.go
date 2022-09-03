package main

import (
	"github.com/championlong/backend-common/design_patterns/dip/di"
	"github.com/championlong/backend-common/design_patterns/dip/ioc"
)

/*
依赖反转原则
控制反转(IOC):“控制”指的是对程序执行流程的控制，而“反转”指的是在没有使用框架之前，程序员自己控制整 个程序的执行。在使用框架之后，整个程序的执行流程可以通过框架来控制。流程的控制权从程序员“反 转”到了框架。
依赖注入(DI):不通过new()的方式在类内部创建依赖类对象，而是 将依赖的类对象在外部创建好之后，通过构造函数、函数参数等方式传递(或注入)给类使用。基于接口而非实现编程。
依赖注入框架(DI Framework):我们只需要通过依赖注入框架提供的扩展点，简单配置 一下所有需要创建的类对象、类与类之间的依赖关系，就可以实现由框架来自动创建对象、管理对象的生命 周期、依赖注入等原本需要程序员来做的事情。
依赖反转原则(DIP):高层模块(high-level modules)不要依赖低层模块(low-level)。 高层模块和低层模块应该通过抽象(abstractions)来互相依赖。除此之外，抽象(abstractions)不要依 赖具体实现细节(details)，具体实现细节(details)依赖抽象(abstractions)。
*/
func main() {
	//控制反转
	application := ioc.NewJunitApplication()
	application.Register(&ioc.UserServiceTest{})
	application.Run()

	//依赖注入
	messageSender := di.SmsSender{}
    di.NewNotification(&messageSender).Send("11111")
}
