package wrap

import (
	"errors"
	"strings"
	"table-export/config"
	"table-export/define"
	"table-export/meta"
)

type sliceWrap struct{}

func (b *sliceWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
		result := "{"
		if origin != "" {
			for _, v := range strSlice {
				content, err := GetOutputValue(exportType, filedType.Value, v)
				if err != nil {
					return nil, err
				}
				valueStr, ok := content.(string)
				if !ok {
					return nil, errors.New("slice get output value error")
				}
				result += valueStr + ","
			}
		}
		result += "}"
		return result, nil
	}
	return nil, nil
}
