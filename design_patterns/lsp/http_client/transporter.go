package http_client

import (
	"fmt"
	"net/http"
)

type Transporter struct {
	HttpClient http.Client
}

func NewTransporter(httpClient http.Client) *Transporter {
	return &Transporter{
		HttpClient: httpClient,
	}
}

func (self *Transporter) SendRequest(request *http.Request) *http.Response {
	response,err := self.HttpClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

func DemoFunction(transporter Transporter){
	response := transporter.SendRequest(&http.Request{})
	fmt.Println(response)
}