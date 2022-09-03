package network

import (
	"io"
	"net/http"
)

type NetworkTransporter struct {
}

var Client = new(http.Client)
func (self *NetworkTransporter) Send(requestUrl, method string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, requestUrl, body)
	if err != nil {
		return nil, err
	}
	resp,err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
