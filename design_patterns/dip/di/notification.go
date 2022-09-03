package di

type MessageSender interface {
	Send (message string)

}

type Notification struct {
	messageSender MessageSender
}

func NewNotification(messageSender MessageSender) *Notification {
	return &Notification{
		messageSender: messageSender,
	}
}

func (self *Notification)Send(message string)  {
	self.messageSender.Send(message)
}