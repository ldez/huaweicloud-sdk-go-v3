/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInstanceStatusResponse struct {
	// 实例ID
	TaskId *string `json:"task_id,omitempty"`
	// 实例创建状态
	TaskStatus *string `json:"task_status,omitempty"`
	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线名字
	PipelineName *string `json:"pipeline_name,omitempty"`
	// 流水线详情页面url
	PipelineUrl *string `json:"pipeline_url,omitempty"`
}

func (o ShowInstanceStatusResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceStatusResponse", string(data)}, " ")
}
