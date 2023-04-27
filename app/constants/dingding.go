package constants

// 钉钉报警机器人
const (
	ALERT_WEBHOOK_ROBOT = "alert-robot"
)

/*
	钉钉审批相关
*/
const (
	DingApplyEventType    = "bpms_instance_change" //审批流类型
	DingApplyFinishStatus = "finish"               //审批完成
	DingApplyAgreeResult  = "agree"                //审批通过

	//系统名称为【Ads广告投放平台(ads.p1staff.com)】的审批模版信息
	DingApplyRoleProcessCode    = ""          //审批模版ID
	DingApplyRoleInfoColumnName = ""          //审批中权限信息列名
	DingApplyRoleManColumnName  = ""          //审批中申请人列名
	DingApplyRoleType           = "dingApply" //通过审批插入类型
)
