package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiViewModelVersionViewCreateDto struct {

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 版本对象ID。
	VersionId string `json:"versionId"`

	// 关系的复制类型。 - BOTH：若存在关系实例引用此数据实例作为源端实例或目标端实例，创建多维版本后的数据实例将继承这些关系实例。 - SOURCE：若存在关系实例引用此数据实例作为源端实例，创建多维版本后的数据实例将继承这些关系实例。 - TARGET：若存在关系实例引用此数据实例作为目标端实例，创建多维版本后的数据实例将继承这些关系实例。 - NONE：创建多维版本后的数据实例将不继承任何关系实例。 - CUSTOM：若指定的关系实体集合对应的关系实例引用此数据实例作为源端实例或目标端实例，创建多维版本后的数据实例将继承这些关系实例。
	WorkCopyType *string `json:"workCopyType,omitempty"`

	// 关系实体名称集合，与workCopyType的值CUSTOM配合使用。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	// 指定不复制的属性，其值将被设置为null。
	NeedSetNull *[]string `json:"needSetNull,omitempty"`

	Item *ObjectReferenceParamDto `json:"item"`
}

func (o MultiViewModelVersionViewCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiViewModelVersionViewCreateDto struct{}"
	}

	return strings.Join([]string{"MultiViewModelVersionViewCreateDto", string(data)}, " ")
}
