package http_client

import "net/http"

type SecurityTransporter struct {
	Transporter
	AppId string
	AppToken string
}

func NewSecurityTransporter(httpClient http.Client, appId, appToken string) *SecurityTransporter {
	return &SecurityTransporter{
		AppId: appId,
		AppToken: appToken,
		Transporter: Transporter{HttpClient: httpClient},
	}
}

func (self *SecurityTransporter)SendRequest(request *http.Request) *http.Response {
	if self.AppId != "" && self.AppToken != ""{

	}
	return self.Transporter.SendRequest(request)
}

