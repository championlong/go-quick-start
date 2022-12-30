package message

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/championlong/backend-common/dingding/config"
	"github.com/championlong/backend-common/dingding/model"
	"github.com/championlong/backend-common/dingding/utils"
	"net/url"
	"strconv"
	"time"
)

var dingdingClient = utils.NewHttpClient()

type MsgType string

const (
	MsgTypeText       MsgType = "text"
	MsgTypeMarkdown   MsgType = "markdown"
	MsgTypeLink       MsgType = "link"
	MsgTypeActionCard MsgType = "actionCard"
	MsgTypeFeedCard   MsgType = "feedCard"
)

type Message struct {
	Msgtype string       `json:"msgtype"`
	Text    *TextMessage `json:"text,omitempty"`
}

type At struct {
	IsAtAll   bool     `json:"isAtAll,omitempty"`   //@所有人
	AtMobiles []string `json:"atMobiles,omitempty"` //@手机号
	AtUsersId []string `json:"atUserIds,omitempty"` //@userid
}

type DingdingMasterJob struct {
	Url           string                     `json:"-"`       //请求url
	KindRobot     string                     `json:"-"`       //机器人种类
	Msgtype       MsgType                    `json:"msgtype"` //发送类型
	Text          *TextMessage               `json:"text,omitempty"`
	Markdown      *MarkdownMessage           `json:"markdown,omitempty"`
	Link          *LinkMessage               `json:"link,omitempty"`
	ActionCard    *ActionCardMessage         `json:"actionCard,omitempty"`
	FeedCard      *FeedCardMessage           `json:"feedCard,omitempty"`
	At            At                         `json:"at,omitempty"`
	CommonMessage interface{}                `json:"-"`
	Query         config.DingdingQueryConfig `json:"-"`
}

type SendMessage interface {
	SendMessage() error
}

func (job *DingdingMasterJob) SendMessage() error {
	switch job.Msgtype {
	case MsgTypeText:
		if value, ok := job.CommonMessage.(TextMessage); ok {
			job.Text = &value
		}
	case MsgTypeMarkdown:
		if value, ok := job.CommonMessage.(MarkdownMessage); ok {
			job.Markdown = &value
		}
	case MsgTypeLink:
		if value, ok := job.CommonMessage.(LinkMessage); ok {
			job.Link = &value
		}
	case MsgTypeActionCard:
		if value, ok := job.CommonMessage.(ActionCardMessage); ok {
			job.ActionCard = &value
		}
	case MsgTypeFeedCard:
		if value, ok := job.CommonMessage.(FeedCardMessage); ok {
			job.FeedCard = &value
		}
	}
	return job.PostDingdingWebHookMsg()
}

func (job *DingdingMasterJob) PostDingdingWebHookMsg() error {
	ctx := fmt.Sprintf("[%s]", job.KindRobot)
	queries := make(map[string]string)
	queries["access_token"] = job.Query.AccessToken
	if job.Query.Encrypt != "" {
		timestampInt := time.Now().UnixNano() / 1e6
		timestamp := strconv.FormatInt(timestampInt, 10)
		secret := job.Query.Encrypt
		signString := timestamp + "\n" + secret
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(signString))
		sha := h.Sum(nil)
		sign := base64.StdEncoding.EncodeToString(sha)

		queries["sign"] = url.QueryEscape(sign)
		queries["timestamp"] = timestamp
	}

	var resBody, err = dingdingClient.PostJson(utils.PackUrl(job.Url, queries), job)
	if err != nil {
		return fmt.Errorf("error while %s. err: %s", ctx, err.Error())
	}

	var data = new(model.DingDingResponse)
	err = json.Unmarshal(resBody, data)
	if err != nil {
		return fmt.Errorf("error while unmarshal %s. err: %s", ctx, err.Error())
	}
	if data.ErrCode != 0 {
		return fmt.Errorf("failed to %s. response: %s", ctx, resBody)
	}
	return nil
}
