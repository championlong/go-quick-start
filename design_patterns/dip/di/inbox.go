package di

import "fmt"

type InboxSender struct {
}

func (self *InboxSender) Send(message string) {
	fmt.Println("inbox message", message)
}
