package system

import (
	"encoding/json"
	"fmt"
	systemModel "github.com/championlong/go-quick-start/internal/app/model/system"
	"github.com/championlong/go-quick-start/internal/app/model/system/request"
	system2 "github.com/championlong/go-quick-start/internal/app/service/system"
	"github.com/championlong/go-quick-start/internal/pkg/constants"
	"github.com/championlong/go-quick-start/internal/pkg/utils"
	"github.com/championlong/go-quick-start/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DingRoles struct{}

// DingApplyRoles 钉钉审批角色回调接口
func (ding *DingRoles) DingApplyRoles(c *gin.Context) {
	var applyRoles request.ApplyRoleRequest
	var eventSubscriptionMsg systemModel.EventSubscriptionMsg
	ctx2 := "[ApplyRoles]"
	if err := c.ShouldBindJSON(&applyRoles); err != nil {
		log.Error("钉钉申请请求解析失败", zap.Error(err))
	}
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")

	callbackCrypto := system2.GetDingTalkCrypto()
	decryptMsg, err := callbackCrypto.GetDecryptMsg(signature, timestamp, nonce, applyRoles.Encrypt)
	if err != nil {
		log.Error("钉钉审批解密请求信息失败", zap.Error(err))
	}
	msg, err := callbackCrypto.GetEncryptMsg("success")
	if err != nil {
		log.Error("钉钉审批获取加密信息失败", zap.Error(err))
	}
	c.JSON(200, msg)
	err = json.Unmarshal([]byte(decryptMsg), &eventSubscriptionMsg)
	if err != nil {
		log.Error("钉钉审批序列化审批信息失败", zap.Error(err))
	}
	if eventSubscriptionMsg.EventType != constants.DingApplyEventType ||
		eventSubscriptionMsg.ProcessCode != constants.DingApplyRoleProcessCode ||
		eventSubscriptionMsg.Status != constants.DingApplyFinishStatus ||
		eventSubscriptionMsg.Result != constants.DingApplyAgreeResult ||
		eventSubscriptionMsg.ProcessInstanceId == "" {
		return
	}
	log.Info(fmt.Sprintf("%s 钉钉审批信息 %s", ctx2, decryptMsg))
	approveInfoResponse, err := system2.GetDingApproveInfo(eventSubscriptionMsg.ProcessInstanceId)
	if err != nil {
		log.Error("获取钉钉详情审批信息失败", zap.Error(err))
		utils.SendDingdingAlertError(fmt.Sprintf("钉钉角色权限获取详细审批失败\n错误信息：%s", err.Error()))
		return
	}
	fmt.Println(approveInfoResponse)
}
