package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ColumnType struct {
	value string
}

type ColumnTypeEnum struct {
	LONG   ColumnType
	STRING ColumnType
	DOUBLE ColumnType
}

func GetColumnTypeEnum() ColumnTypeEnum {
	return ColumnTypeEnum{
		LONG: ColumnType{
			value: "Long",
		},
		STRING: ColumnType{
			value: "String",
		},
		DOUBLE: ColumnType{
			value: "Double",
		},
	}
}

func (c ColumnType) Value() string {
	return c.value
}

func (c ColumnType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ColumnType) UnmarshalJSON(b []byte) error {
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
