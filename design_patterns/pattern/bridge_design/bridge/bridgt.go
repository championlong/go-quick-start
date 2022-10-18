package bridge

type IMsgSender interface {
	Send(msg string) error
}

type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func (s *EmailMsgSender) Send(msg string) error {
	return nil
}

type WechatMsgSender struct {
	emails []string
}

func NewWechatMsgSender(emails []string) *WechatMsgSender {
	return &WechatMsgSender{emails: emails}
}

func (s *WechatMsgSender) Send(msg string) error {
	return nil
}


type INotification interface {
	Notify(msg string) error
}

type SevereNotification struct {
	sender IMsgSender
}

func NewSevereNotification(sender IMsgSender) *SevereNotification {
	return &SevereNotification{sender: sender}
}

func (n *SevereNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
