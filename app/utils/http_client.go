package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	juju_ratelimit "github.com/juju/ratelimit"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

type HttpClient struct {
	Header        http.Header
	Client        *http.Client
	RateLimiter   map[string]*juju_ratelimit.Bucket
	RetryTimes    int //重试次数
	RetryInterval time.Duration
}

const (
	DEFAULT_USER_AGENT = ""
)

var DefaultRateLimiter = map[string]*juju_ratelimit.Bucket{
	DEFAULT_USER_AGENT: juju_ratelimit.NewBucket(time.Second, 1), //1 per second
}

func NewHttpClient(rateLimiter map[string]*juju_ratelimit.Bucket, retryTimes int, timeout time.Duration) HttpClient {
	jar, _ := cookiejar.New(nil)
	return HttpClient{
		Header:      make(http.Header),
		Client:      &http.Client{Jar: jar},
		RateLimiter: rateLimiter,
		RetryTimes:  retryTimes,
	}
}

func (self *HttpClient) SetRetryTimes(n int) {
	self.RetryTimes = n
}

func (self *HttpClient) SetRetryInterval(i time.Duration) *HttpClient {
	self.RetryInterval = i
	return self
}

//fillInterval: 多久填充1次
//capacity: 1次填充多少
func (self *HttpClient) SetBucket(bucketName string, fillInterval time.Duration, capacity int64) {
	self.RateLimiter[bucketName] = juju_ratelimit.NewBucketWithQuantum(fillInterval, capacity, capacity)
}

//make a http get request
func (self *HttpClient) Get(requestUrl string) (*http.Response, error) {
	return self.GetWithCookies(requestUrl, nil)
}

//make a http get request with cookies
func (self *HttpClient) GetWithCookies(requestUrl string, cookies []*http.Cookie) (*http.Response, error) {
	fmt.Printf("%s %s\n", http.MethodGet, requestUrl)

	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("nil url")
	}

	q := u.Query()
	u.RawQuery = q.Encode() //urlencode

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	self.copyRequestHeader(req)
	if cookies != nil {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}

	return self.Client.Do(req)
}

//make a http get request, and return the string body
func (self *HttpClient) GetBody(requestUrl string) ([]byte, error) {
	fmt.Printf("%s %s\n", http.MethodGet, requestUrl)

	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	self.copyRequestHeader(req)

	resp, err := self.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer self.closeRespBody(requestUrl, resp)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

//make a http post request, and return the []byte body
func (self *HttpClient) Post(requestUrl, contentType string, postBody io.Reader) ([]byte, error) {
	fmt.Printf(http.MethodPost + " " + requestUrl + "\n")

	if !strings.Contains(contentType, "multipart/form-data") {
		fmt.Printf(http.MethodPost+" params: %+v\n", postBody)
	}

	//埋点
	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode

	req, err := http.NewRequest(http.MethodPost, requestUrl, postBody)
	if err != nil {
		return nil, err
	}

	self.copyRequestHeader(req)
	req.Header.Set("Content-Type", contentType)

	resp, err := self.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer self.closeRespBody(requestUrl, resp)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func (self *HttpClient) PostForm(requestUrl string, data url.Values) ([]byte, error) {
	return self.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func (self *HttpClient) PostJson(requestUrl string, data interface{}) ([]byte, error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bytesData)
	return self.Post(requestUrl, "application/json;charset=UTF-8", reader)
}

//post json，带重试
func (self *HttpClient) PostJsonWithRetry(requestUrl string, data interface{}) ([]byte, error) {
	var err = errors.New("")
	var result []byte
	for i := 0; i < self.RetryTimes && err != nil; i++ {
		result, err = self.PostJson(requestUrl, data)
	}
	return result, err
}

//上传
func (self *HttpClient) Upload(requestUrl, fieldName, filePath string, params map[string]string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Printf("http_client.upload: failed to close file %\n", filePath)
		}
	}()

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	defer func() {
		if err = writer.Close(); err != nil {
			fmt.Printf("http_client.upload: failed to close write for file %s\n", filePath)
		}
	}()

	part, err := writer.CreateFormFile(fieldName, filePath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		if err = writer.WriteField(key, val); err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(http.MethodPost, requestUrl, buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := self.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer self.closeRespBody(requestUrl, resp)

	return ioutil.ReadAll(resp.Body)
}

func (self *HttpClient) SetHttpProxy(httpProxy string) {
	proxy, _ := url.Parse(httpProxy)
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	self.Client.Transport = tr
}

func (self *HttpClient) SetTimeout(timeout time.Duration) {
	self.Client.Timeout = timeout
}

//从指定链接下载文件到本地
func (self *HttpClient) DownloadFile(requestUrl, path string) (err error) {
	fileBody, err := self.GetBody(requestUrl)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, fileBody, 0644)
}

//控制是否需要重定向，因httpclient重定向后不会带cookie，所以有时需要手动重定向
func (self *HttpClient) AutoRedirect(is bool) {
	if is {
		self.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return nil
		}
	} else {
		self.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
}

//set self.Header into request.Header
func (self *HttpClient) copyRequestHeader(req *http.Request) {
	if self.Header != nil {
		for k, values := range self.Header {
			if len(values) == 0 {
				continue
			}
			if len(values) == 1 {
				req.Header.Set(k, values[0])
			} else {
				req.Header.Del(k)
				for _, value := range values {
					req.Header.Add(k, value)
				}
			}
		}
	}
	req.Header.Set("User-Agent", DEFAULT_USER_AGENT)
}

//关闭http body
func (self *HttpClient) closeRespBody(requestUrl string, resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}
	if err := resp.Body.Close(); err != nil {
		fmt.Printf("http_client.close_resp_body: failed to close %s. err: %s", requestUrl, err.Error())
	}
}

//消耗令牌桶的令牌
func (self *HttpClient) consumeBucket(label string) {
	if self.RateLimiter != nil && self.RateLimiter[label] != nil {
		self.RateLimiter[label].Wait(1)
	}
}
