package domain

import "fmt"

type Notification struct {
}

func (self *Notification) Notify(message string) {
	fmt.Println(message)
}