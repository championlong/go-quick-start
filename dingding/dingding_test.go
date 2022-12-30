package dingding

import (
	"fmt"
	"github.com/championlong/backend-common/dingding/config"
	"github.com/championlong/backend-common/dingding/message"
	"testing"
)

func TestSend(t *testing.T) {
	query := make(map[string]config.DingdingQueryConfig)
	query["rebort"] = config.DingdingQueryConfig{
		Encrypt:     "SEC28c01255689eed0515a1a2032daafe372b3c86667428bc8a2035f9ef3d0174fc",
		AccessToken: "8092dbdace9b4fc2ef20fc0b29c2031d7174a6ef7a036941005625d1b80a0062",
	}
	dingdingConfig = &config.DingdingConfig{
		Url: "https://oapi.dingtalk.com/robot/send",
		DingdingQuery: query,
	}
	Init(*dingdingConfig)
	err := SendMessage("rebort", message.MsgTypeText, message.TextMessage{
		Content: "text",
	}, message.At{})
	fmt.Println(err)
}
