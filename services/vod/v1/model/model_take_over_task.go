package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TakeOverTask struct {
	// 桶 <br/>

	Bucket *string `json:"bucket,omitempty"`
	// 目录

	Object *string `json:"object,omitempty"`
	// 托管类型：0表示存储到点播桶 1表示存储在租户桶  2表示存储到租户桶，并且源文件名跟随<br/>

	HostType *int32 `json:"host_type,omitempty"`
	// 输出桶 <br/>

	OutputBucket *string `json:"output_bucket,omitempty"`
	// 输出路径 <br/>

	OutputPath *string `json:"output_path,omitempty"`
	// 任务ID <br/>

	TaskId *string `json:"task_id,omitempty"`

	Suffix *[]string `json:"suffix,omitempty"`
	// 转码模板组 <br/>

	TemplateGroupName *string `json:"template_group_name,omitempty"`
	// 创建时间<br/>

	CreateTime *string `json:"create_time,omitempty"`
	// 结束时间<br/>

	EndTime *string `json:"end_time,omitempty"`

	Status *TakeOverTaskStatus `json:"status,omitempty"`
	// 媒资的任务执行描述汇总。

	ExecDesc *string `json:"exec_desc,omitempty"`
}

func (o TakeOverTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TakeOverTask struct{}"
	}

	return strings.Join([]string{"TakeOverTask", string(data)}, " ")
}

type TakeOverTaskStatus struct {
	value string
}

type TakeOverTaskStatusEnum struct {
	PROCESSING TakeOverTaskStatus
	SUCCEED    TakeOverTaskStatus
	FAILED     TakeOverTaskStatus
}

func GetTakeOverTaskStatusEnum() TakeOverTaskStatusEnum {
	return TakeOverTaskStatusEnum{
		PROCESSING: TakeOverTaskStatus{
			value: "PROCESSING",
		},
		SUCCEED: TakeOverTaskStatus{
			value: "SUCCEED",
		},
		FAILED: TakeOverTaskStatus{
			value: "FAILED",
		},
	}
}

func (c TakeOverTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TakeOverTaskStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
