package request

import (
	"github.com/championlong/go-quick-start/internal/app/model/common/request"
	"github.com/championlong/go-quick-start/internal/app/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
