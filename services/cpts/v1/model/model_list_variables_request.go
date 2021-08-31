package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListVariablesRequest struct {
	// 变量类型

	VariableType int32 `json:"variable_type"`
	// 测试工程id

	TestSuiteId int32 `json:"test_suite_id"`
}

func (o ListVariablesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVariablesRequest struct{}"
	}

	return strings.Join([]string{"ListVariablesRequest", string(data)}, " ")
}
