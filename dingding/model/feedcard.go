package model

type FeedCardMessage struct {
	Links []Links `json:"links"`
}

type Links struct {
	Title      string `json:"title"`      //单条信息文本
	MessageURL string `json:"messageURL"` //点击单条信息到跳转链接
	PicURL     string `json:"picURL"`     //单条信息后面图片的URL
}

