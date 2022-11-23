package message

import "github.com/championlong/backend-common/dingding/model"

type MsgType string

const (
	MsgTypeText       MsgType = "text"
	MsgTypeMarkdown   MsgType = "markdown"
	MsgTypeLink       MsgType = "link"
	MsgTypeActionCard MsgType = "actionCard"
	MsgTypeFeedCard   MsgType = "feedCard"
)

type SendMessage interface {
	GetMessage() *Message
}

type Message struct {
	Msgtype string       `json:"msgtype"`
	Text    *TextMessage `json:"text,omitempty"`
}

type at struct {
	IsAtAll   bool     `json:"isAtAll,omitempty"`   //@所有人
	AtMobiles []string `json:"atMobiles,omitempty"` //@手机号
	AtUsersId []string `json:"atUserIds,omitempty"` //@userid
}

type DingdingMasterJob struct {
	Url        string                    //请求url
	KindRobot  string                    //机器人种类
	Msgtype    MsgType                   `json:"msgtype"` //发送类型
	Text       *TextMessage              `json:"text,omitempty"`
	Markdown   *MarkdownMessage          `json:"markdown,omitempty"`
	Link       *LinkMessage              `json:"link,omitempty"`
	ActionCard *ActionCardMessage        `json:"actionCard,omitempty"`
	FeedCard   *FeedCardMessage          `json:"feedCard,omitempty"`
	At         at                        `json:"at,omitempty"`
	Query      model.DingdingQueryConfig `json:"-"`
}
