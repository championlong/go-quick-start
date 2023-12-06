package constants

/*
	钉钉报警机器人相关
*/

const (
	ALERT_WEBHOOK_ROBOT = "alert-robot"
)

/*
	钉钉审批相关
*/
const (
	DingApplyEventType    = "bpms_instance_change" // 审批流类型
	DingApplyFinishStatus = "finish"               // 审批完成
	DingApplyAgreeResult  = "agree"                // 审批通过

	DingApplyRoleProcessCode = "" // 审批模版ID
)
