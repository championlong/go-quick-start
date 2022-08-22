package main

import (
	"github.com/championlong/backend-common/design_patterns/lsp/http_client"
	"net/http"
)
/*
里约替换
子类对象(object of subtype/derived class)能够替换程序(program)中父类对象(object of base/parent class)出现的任何地方，
并且保证 原来程序的逻辑行为(behavior)不变及正确性不被破坏。
 */
func main() {
	http_client.NewSecurityTransporter(http.Client{},"","").SendRequest(&http.Request{})
}
