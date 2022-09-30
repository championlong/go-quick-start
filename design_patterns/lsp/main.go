package main

import (
	"github.com/championlong/backend-common/design_patterns/lsp/http_client"
	"net/http"
)
/*
里约替换
子类对象(object of subtype/derived class)能够替换程序(program)中父类对象(object of base/parent class)出现的任何地方，并且保证原来程序的逻辑行为(behavior)不变及正确性不被破坏。
父类定义了函数的“约定”(或者叫协议)，那子类可以改变函数的内部实现逻辑，但不能改变函数的原有“约定”。这里的“约定”包括:函数声明要实现 的功能;对输入、输出、异常的约定;甚至包括注释中所罗列的任何特殊说明。
*/
func main() {
	http_client.NewSecurityTransporter(http.Client{},"","").SendRequest(&http.Request{})
}
