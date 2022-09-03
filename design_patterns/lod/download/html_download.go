package download

import (
	"github.com/championlong/backend-common/design_patterns/lod/network"
	"net/http"
)

type HtmlDownloader struct {
	transporter network.NetworkTransporter
}

func (self *HtmlDownloader) downloadHtml(requestUrl string) (string, error) {
	body, err := self.transporter.Send(requestUrl, http.MethodGet, nil)
	return string(body), err
}
