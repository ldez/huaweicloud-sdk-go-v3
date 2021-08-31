package model

import (
	"encoding/json"

	"strings"
)

type MysqlVolume struct {
	// 磁盘大小。默认值为40，单位GB。 取值范围：40~128000，必须为10的整数倍。

	Size string `json:"size"`
}

func (o MysqlVolume) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlVolume struct{}"
	}

	return strings.Join([]string{"MysqlVolume", string(data)}, " ")
}
