package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateWatermarkTemplateReq struct {
	// 水印模板名称<br/>

	Name *string `json:"name,omitempty"`
	// 水印类型，当前只支持Image（图片水印）<br/>

	WatermarkType *CreateWatermarkTemplateReqWatermarkType `json:"watermark_type,omitempty"`
	// type设置为Image时有效，目前包括Original（只做简单缩放，不做其他处理），Transparent（图片底色透明），Grayed（彩色图片变灰）<br/>

	ImageProcess *CreateWatermarkTemplateReqImageProcess `json:"image_process,omitempty"`
	// 水印图片相对输出视频的水平偏移量，默认值是0<br/>

	Dx *string `json:"dx,omitempty"`
	// 水印图片相对输出视频的垂直偏移量，默认值是0<br/>

	Dy *string `json:"dy,omitempty"`
	// 水印的位置<br/>

	Position *CreateWatermarkTemplateReqPosition `json:"position,omitempty"`
	// 水印图片宽<br/>

	Width *string `json:"width,omitempty"`
	// 水印图片高<br/>

	Height *string `json:"height,omitempty"`
	// 水印开始时间<br/>

	TimelineStart *string `json:"timeline_start,omitempty"`
	// 水印持续时间<br/>

	TimelineDuration *string `json:"timeline_duration,omitempty"`
	// 水印图片格式类型<br/>

	Type *string `json:"type,omitempty"`
	// 水印图片MD5值<br/>

	Md5 *string `json:"md5,omitempty"`
}

func (o CreateWatermarkTemplateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWatermarkTemplateReq struct{}"
	}

	return strings.Join([]string{"CreateWatermarkTemplateReq", string(data)}, " ")
}

type CreateWatermarkTemplateReqWatermarkType struct {
	value string
}

type CreateWatermarkTemplateReqWatermarkTypeEnum struct {
	IMAGE CreateWatermarkTemplateReqWatermarkType
	TEXT  CreateWatermarkTemplateReqWatermarkType
}

func GetCreateWatermarkTemplateReqWatermarkTypeEnum() CreateWatermarkTemplateReqWatermarkTypeEnum {
	return CreateWatermarkTemplateReqWatermarkTypeEnum{
		IMAGE: CreateWatermarkTemplateReqWatermarkType{
			value: "IMAGE",
		},
		TEXT: CreateWatermarkTemplateReqWatermarkType{
			value: "TEXT",
		},
	}
}

func (c CreateWatermarkTemplateReqWatermarkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateWatermarkTemplateReqWatermarkType) UnmarshalJSON(b []byte) error {
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

type CreateWatermarkTemplateReqImageProcess struct {
	value string
}

type CreateWatermarkTemplateReqImageProcessEnum struct {
	ORIGINAL    CreateWatermarkTemplateReqImageProcess
	TRANSPARENT CreateWatermarkTemplateReqImageProcess
	GRAYED      CreateWatermarkTemplateReqImageProcess
}

func GetCreateWatermarkTemplateReqImageProcessEnum() CreateWatermarkTemplateReqImageProcessEnum {
	return CreateWatermarkTemplateReqImageProcessEnum{
		ORIGINAL: CreateWatermarkTemplateReqImageProcess{
			value: "ORIGINAL",
		},
		TRANSPARENT: CreateWatermarkTemplateReqImageProcess{
			value: "TRANSPARENT",
		},
		GRAYED: CreateWatermarkTemplateReqImageProcess{
			value: "GRAYED",
		},
	}
}

func (c CreateWatermarkTemplateReqImageProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateWatermarkTemplateReqImageProcess) UnmarshalJSON(b []byte) error {
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

type CreateWatermarkTemplateReqPosition struct {
	value string
}

type CreateWatermarkTemplateReqPositionEnum struct {
	TOPRIGHT    CreateWatermarkTemplateReqPosition
	TOPLEFT     CreateWatermarkTemplateReqPosition
	BOTTOMRIGHT CreateWatermarkTemplateReqPosition
	BOTTOMLEFT  CreateWatermarkTemplateReqPosition
}

func GetCreateWatermarkTemplateReqPositionEnum() CreateWatermarkTemplateReqPositionEnum {
	return CreateWatermarkTemplateReqPositionEnum{
		TOPRIGHT: CreateWatermarkTemplateReqPosition{
			value: "TOPRIGHT",
		},
		TOPLEFT: CreateWatermarkTemplateReqPosition{
			value: "TOPLEFT",
		},
		BOTTOMRIGHT: CreateWatermarkTemplateReqPosition{
			value: "BOTTOMRIGHT",
		},
		BOTTOMLEFT: CreateWatermarkTemplateReqPosition{
			value: "BOTTOMLEFT",
		},
	}
}

func (c CreateWatermarkTemplateReqPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateWatermarkTemplateReqPosition) UnmarshalJSON(b []byte) error {
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
