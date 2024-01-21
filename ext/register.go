package ext

import (
	"github.com/821869798/table-export/field_type"
	"strings"
)

var (
	extFieldTypes = make(map[string]field_type.IExtFieldType)
)

func RegisterExtFieldType(extFieldType field_type.IExtFieldType) {
	extFieldTypes[strings.ToLower(extFieldType.Name())] = extFieldType
}

// GetExistExtFieldType 获取已有类型的扩展字段，方便嵌套，例如一个类包含了另一个类
func GetExistExtFieldType(name string) (field_type.IExtFieldType, bool) {
	v, ok := extFieldTypes[strings.ToLower(name)]
	return v, ok
}
