package main

import "github.com/championlong/backend-common/design_patterns/dip/ioc"

func main() {
	ioc.NewJunitApplication().Register(&ioc.UserServiceTest{})
}
