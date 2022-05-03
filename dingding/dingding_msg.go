package dingding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/championlong/backend-common/dingding/model"
	"github.com/championlong/backend-common/dingding/utils"
	"net/url"
	"strconv"
	"time"
)

const (
	//机器人发送消息类型
	WEBHOOK_SEND_TEXT     = "text"
	WEBHOOK_SEND_MARKDOWN = "markdown"
)

type DingdingMasterJob struct {
	Url                string   //请求url
	KindRobot          string   //机器人种类
	Content            string   //消息内容
	IsAtAll            bool     //@所有人
	AtMobiles          []string //@手机号
	AtUsersId          []string //@userid
	Type               string   //发送类型
	Title              string   //标题
	BtnOrientation     string   //按钮排列
	ActionCardBtnValue []model.ActionCardBtn
	Query              model.DingdingQueryConfig
	ConfigType         interface{}
}

type DingDingJobAdapter struct {
}

var DingDingSend = &DingDingJobAdapter{}
var dingdingClient = utils.NewHttpClient()
var DingdingDynamicQuery map[string]model.DingdingQueryConfig

// GetConfigWebHookRobot 发送群组机器人通知，走钉钉官网api
func (t *DingDingJobAdapter) GetConfigWebHookRobot(kindRobot string) DingDingSendJob {
	dingdingJob := new(DingdingMasterJob)
	dingdingJob.KindRobot = kindRobot
	dingdingJob.Url = GetConfig().WebHook.Url
	notify := GetConfig().WebHook.DingdingQuery
	dingdingJob.Query = notify[kindRobot]
	return dingdingJob
}

// GetDynamicWebHookRobot 后续支持动态添加的钉钉机器人
func (t *DingDingJobAdapter) GetDynamicWebHookRobot(kindRobot string) DingDingSendJob {
	dingdingJob := new(DingdingMasterJob)
	dingdingJob.KindRobot = kindRobot
	dingdingJob.Url = GetConfig().WebHook.Url
	query := DingdingDynamicQuery[kindRobot]
	dingdingJob.Query = query
	return dingdingJob
}

type DingDingSendJob interface {
	DingDingSendGroupText(content string, atMobiles, atUsersId []string, isAtAll bool) error                               //钉钉群组发text类型信息
	DingDingSendMarkdown(content, title string, atMobiles, atUsersId []string, isAtAll bool) error                         //钉钉群组发markdown类型信息 	//钉钉群组发送工作通知Markdown类型，走公司api
}

//conetne:消息内容
//atMobiles:@手机号
//atUsersId:@用户ID
//isAtAll:@所有人
//指定机器人向钉钉群组发送text消息
func (dingdingJob *DingdingMasterJob) DingDingSendGroupText(content string, atMobiles, atUsersId []string, isAtAll bool) error {
	dingdingJob.Content = content
	dingdingJob.AtMobiles = atMobiles
	dingdingJob.AtUsersId = atUsersId
	dingdingJob.IsAtAll = isAtAll
	dingdingJob.Type = WEBHOOK_SEND_TEXT

	err := dingdingJob.PostDingdingWebHookMsg()
	if err != nil {
		utils.Err("DingDingSendGroupText send fail.robot:%s content:%s err:%s", dingdingJob.KindRobot, content, err.Error())
		return err
	}
	utils.Info("send dingding successfully %s", dingdingJob.KindRobot)
	return nil
}

// DingDingSendMarkdown 指定机器人向钉钉群组发送Markdown消息
func (dingdingJob *DingdingMasterJob) DingDingSendMarkdown(content, title string, atMobiles, atUsersId []string, isAtAll bool) error {
	dingdingJob.Content = content
	dingdingJob.Title = title
	dingdingJob.AtMobiles = atMobiles
	dingdingJob.AtUsersId = atUsersId
	dingdingJob.IsAtAll = isAtAll
	dingdingJob.Type = WEBHOOK_SEND_MARKDOWN

	err := dingdingJob.PostDingdingWebHookMsg()
	if err != nil {
		utils.Err("DingDingSendMarkdown send fail.robot:%s content:%s err:%s", dingdingJob.KindRobot, content, err.Error())
		return err
	}
	utils.Info("send dingding successfully %s", dingdingJob.KindRobot)
	return nil
}

func (dingdingJob *DingdingMasterJob) PostDingdingWebHookMsg() error {
	ctx := fmt.Sprintf("%s. content: %s", dingdingJob.KindRobot, dingdingJob.Content)
	queries := make(map[string]string)
	msgInfo := make(map[string]interface{})
	queries["access_token"] = dingdingJob.Query.AccessToken
	if dingdingJob.Query.IsEncrypt {
		timestampInt := time.Now().UnixNano() / 1e6
		timestamp := strconv.FormatInt(timestampInt, 10)
		secret := dingdingJob.Query.Encrypt
		signString := timestamp + "\n" + secret
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(signString))
		sha := h.Sum(nil)
		sign := base64.StdEncoding.EncodeToString(sha)

		queries["sign"] = url.QueryEscape(sign)
		queries["timestamp"] = timestamp
	}
	if dingdingJob.Query.IsKeysWord {
		dingdingJob.Content = fmt.Sprintf("%s\n%s", dingdingJob.Query.KeysWord, dingdingJob.Content)
	}

	switch dingdingJob.Type {
	case WEBHOOK_SEND_TEXT:
		msgInfo = dingdingJob.getTextMsg()
	case WEBHOOK_SEND_MARKDOWN:
		msgInfo = dingdingJob.getMarkdownMsg()
	}
	var resBody, err = dingdingClient.PostJsonWithRetry(utils.PackUrl(dingdingJob.Url, queries), msgInfo)

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

func (dingdingJob *DingdingMasterJob) getTextMsg() map[string]interface{} {
	atInfo := make(map[string]interface{})
	var subText = map[string]string{
		"content": dingdingJob.Content,
	}
	if len(dingdingJob.AtMobiles) > 0 && dingdingJob.AtMobiles[0] != "" {
		atInfo["atMobiles"] = dingdingJob.AtMobiles
	}
	if len(dingdingJob.AtUsersId) > 0 && dingdingJob.AtUsersId[0] != "" {
		atInfo["atUserIds"] = dingdingJob.AtUsersId
	}
	if dingdingJob.IsAtAll {
		atInfo["isAtAll"] = dingdingJob.IsAtAll
	}
	var text = map[string]interface{}{
		"msgtype": "text",
		"text":    subText,
		"at":      atInfo,
	}
	return text
}

func (dingdingJob *DingdingMasterJob) getMarkdownMsg() map[string]interface{} {
	atInfo := make(map[string]interface{})
	var subMarkdown = map[string]string{
		"title": dingdingJob.Title,
		"text":  dingdingJob.Content,
	}
	if len(dingdingJob.AtMobiles) > 0 && dingdingJob.AtMobiles[0] != "" {
		atInfo["atMobiles"] = dingdingJob.AtMobiles
	}
	if len(dingdingJob.AtUsersId) > 0 && dingdingJob.AtUsersId[0] != "" {
		atInfo["atUserIds"] = dingdingJob.AtUsersId
	}
	if dingdingJob.IsAtAll {
		atInfo["isAtAll"] = dingdingJob.IsAtAll
	}
	var text = map[string]interface{}{
		"msgtype":  "markdown",
		"markdown": subMarkdown,
		"at":       atInfo,
	}
	return text
}

func (dingdingJob *DingdingMasterJob) getActionCardMsg() map[string]interface{} {
	var subActionCard = map[string]interface{}{
		"title":          dingdingJob.Title,
		"text":           dingdingJob.Content,
		"btnOrientation": dingdingJob.BtnOrientation,
		"btns":           dingdingJob.ActionCardBtnValue,
	}
	var text = map[string]interface{}{
		"msgtype":    "actionCard",
		"actionCard": subActionCard,
	}
	return text
}
