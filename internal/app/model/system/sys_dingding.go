package system

type AccessTokenResponse struct {
	Errcode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	ErrMsg      string `json:"errmsg"`
	ExpiresIn   int    `json:"expires_in"`
}

type ApproveInfoResponse struct {
	RequestId       string               `json:"request_id"`
	Errcode         int                  `json:"errcode"`
	Errmsg          string               `json:"errmsg"`
	ProcessInstance ProcessInstanceTopVo `json:"process_instance"`
}

type ApproveInfoRequest struct {
	ProcessInstanceId string `json:"process_instance_id"`
}

type ProcessInstanceTopVo struct {
	Title                      string                 `json:"title"`
	CreateTime                 string                 `json:"create_time"`
	FinishTime                 string                 `json:"finish_time"`
	OriginatorUserid           string                 `json:"originator_userid"`
	OriginatorDeptId           string                 `json:"originator_dept_id"`
	Status                     string                 `json:"status"`
	ApproverUserids            string                 `json:"approver_userids"`
	CcUserids                  []string               `json:"cc_userids"`
	Result                     string                 `json:"result"`
	BusinessId                 string                 `json:"business_id"`
	OperationRecords           []OperationRecordsVo   `json:"operation_records"`
	Tasks                      []TaskTopVo            `json:"tasks"`
	OriginatorDeptName         string                 `json:"originator_dept_name"`
	BizAction                  string                 `json:"biz_action"`
	AttachedProcessInstanceIds []string               `json:"attached_process_instance_ids"`
	FormComponentValues        []FormComponentValueVo `json:"form_component_values"`
	MainProcessInstanceId      string                 `json:"main_process_instance_id"`
}

type OperationRecordsVo struct {
	Userid          string `json:"userid"`
	Date            string `json:"date"`
	OperationType   string `json:"operation_type"`
	OperationResult string `json:"operation_result"`
	Remark          string `json:"remark"`
}

type TaskTopVo struct {
	Userid     string `json:"userid"`
	TaskStatus string `json:"task_status"`
	TaskResult string `json:"task_result"`
	CreateTime string `json:"create_time"`
	FinishTime string `json:"finish_time"`
	TaskId     string `json:"task_id"`
	Url        string `json:"url"`
}

type FormComponentValueVo struct {
	Name          string `json:"name"`
	Value         string `json:"value"`
	ExtValue      string `json:"ext_value"`
	Id            string `json:"id"`
	ComponentType string `json:"component_type"`
}

type EventSubscriptionMsg struct {
	ProcessInstanceId string `json:"processInstanceId"`
	FinishTime        int64  `json:"finishTime"`
	CorpId            string `json:"corpId"`
	EventType         string `json:"EventType"`
	BusinessId        string `json:"businessId"`
	Title             string `json:"title"`
	Status            string `json:"type"`
	Url               string `json:"url"`
	Result            string `json:"result"`
	CreateTime        int64  `json:"createTime"`
	ProcessCode       string `json:"processCode"`
	BizCategoryId     string `json:"bizCategoryId"`
	BusinessType      string `json:"businessType"`
	StaffId           string `json:"staffId"`
}
