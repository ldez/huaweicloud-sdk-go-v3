package model

import (
	"encoding/json"

	"strings"
)

// 工作项属性
type IssueResponseV4 struct {
	// 实际工时

	ActualWorkHours *float64 `json:"actual_work_hours,omitempty"`
	// 抄送人

	AssignedCcUser *[]IssueUser `json:"assigned_cc_user,omitempty"`

	AssignedUser *IssueUser `json:"assigned_user,omitempty"`
	// 预计开始时间，年-月-日

	BeginTime *string `json:"begin_time,omitempty"`
	// 创建时间 年-月-日 时:分:秒

	CreatedTime *string `json:"created_time,omitempty"`

	Creator *IssueUser `json:"creator,omitempty"`
	// 自定义属性值

	CustomFields *[]CustomField `json:"custom_fields,omitempty"`

	Developer *IssueUser `json:"developer,omitempty"`

	Domain *IssueItemSfV4Domain `json:"domain,omitempty"`
	// 工作项进度值

	DoneRatio *int32 `json:"done_ratio,omitempty"`
	// 预计结束时间，年-月-日

	EndTime *string `json:"end_time,omitempty"`
	// 预计工时

	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty"`
	// 工作项项id

	Id *int32 `json:"id,omitempty"`

	Project *IssueProjectResponseV4 `json:"project,omitempty"`

	Iteration *IssueItemSfV4Iteration `json:"iteration,omitempty"`

	Module *IssueItemSfV4Module `json:"module,omitempty"`
	// 标题

	Name *string `json:"name,omitempty"`

	ParentIssue *CreateIssueResponseV4ParentIssue `json:"parent_issue,omitempty"`

	Priority *IssueItemSfV4Priority `json:"priority,omitempty"`

	Severity *IssueItemSfV4Severity `json:"severity,omitempty"`

	Status *IssueItemSfV4Status `json:"status,omitempty"`

	Tracker *IssueItemSfV4Tracker `json:"tracker,omitempty"`
	// 更新时间 年-月-日 时:分:秒

	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o IssueResponseV4) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IssueResponseV4 struct{}"
	}

	return strings.Join([]string{"IssueResponseV4", string(data)}, " ")
}
