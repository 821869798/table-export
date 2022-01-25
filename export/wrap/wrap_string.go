package wrap

import (
	"strings"
	"table-export/export/common"
	"table-export/meta"
)

type stringWrap struct{}

func (b *stringWrap) OutputValue(exportType common.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case common.ExportType_Lua:
		newValue := strings.Replace(origin, "\n", "\\n", -1)
		newValue = strings.Replace(newValue, "\"", "\\\"", -1)
		newValue = "\"" + newValue + "\""
		return newValue, nil
	}
	return nil, nil
}
