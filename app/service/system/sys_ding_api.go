package system

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/championlong/go-quick-start/app/constants"
	"github.com/championlong/go-quick-start/app/global"
	"github.com/championlong/go-quick-start/app/model/system"
	"github.com/championlong/go-quick-start/app/utils"
	"go.uber.org/zap"
	"net/url"
	"time"
)

var client = utils.HttpClient{}

const (
	appKey            = ""
	appSecret         = ""
	getAccessTokenUrl = "https://oapi.dingtalk.com/gettoken"
	getApproveInfoUrl = "https://oapi.dingtalk.com/topapi/processinstance/get"
)

func getDingAccessToken(ctx context.Context, redisKey string) (string, error) {
	accessToken := global.GVA_REDIS.Get(ctx, redisKey).String()

	if accessToken != "" {
		return accessToken, nil
	}

	var accessTokenResponse system.AccessTokenResponse
	params := url.Values{}
	accessTokenUrl, err := url.ParseRequestURI(getAccessTokenUrl)
	if err != nil {
		global.GVA_LOG.Error("GetDingdingAccessToken illegal URI!", zap.Error(err))
		return token, err
	}
	params.Set("appkey", appKey)
	params.Set("appsecret", appSecret)
	accessTokenUrl.RawQuery = params.Encode()
	body, err := client.GetBody(accessTokenUrl.String())
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return "", err
	}
	if accessTokenResponse.ErrMsg != "ok" {
		return "", err
	}
	if accessTokenResponse.AccessToken != "" {
		err = global.GVA_REDIS.Set(ctx, redisKey, accessTokenResponse.AccessToken, time.Minute*90).Err()
		if err != nil {
			global.GVA_LOG.Error("GetDingdingAccessToken set redis fail", zap.Error(err))
		}
		//redis存储失败不影响正常返回值
		return accessTokenResponse.AccessToken, nil
	}
	return "", err
}

func GetDingApproveInfo(processInstanceId string) (approveInfoResponse system.ApproveInfoResponse, err error) {
	accessToken, err := getDingAccessToken(context.Background(), constants.DINGDING_ACCESS_TOKEN_ROLE)
	if err != nil {
		return approveInfoResponse, err
	}
	if accessToken == "" {
		return approveInfoResponse, fmt.Errorf(fmt.Sprintf("dingding accessToken is null"))
	}

	params := url.Values{}
	accessTokenUrl, err := url.ParseRequestURI(getApproveInfoUrl)
	if err != nil {
		return approveInfoResponse, err
	}
	params.Set("access_token", accessToken)
	accessTokenUrl.RawQuery = params.Encode()
	var approveInfoRequest system.ApproveInfoRequest
	approveInfoRequest.ProcessInstanceId = processInstanceId

	body, err := client.PostJson(accessTokenUrl.String(), approveInfoRequest)
	if err != nil {
		return approveInfoResponse, err
	}
	err = json.Unmarshal(body, &approveInfoResponse)
	if err != nil {
		return approveInfoResponse, err
	}
	if approveInfoResponse.Errmsg != "ok" {
		return approveInfoResponse, fmt.Errorf(fmt.Sprintf("GetDingdingApproveInfo fail"))
	}

	return approveInfoResponse, err
}