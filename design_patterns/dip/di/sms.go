package di

import "fmt"

type SmsSender struct {
}

func (self *SmsSender) Send(message string) {
	fmt.Println("sms message", message)
}
