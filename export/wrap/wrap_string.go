package wrap

import (
	"errors"
	"strings"
	"table-export/define"
	"table-export/meta"
)

type stringWrap struct{}

func (b *stringWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		newValue := strings.Replace(origin, "\\", "\\\\", -1)
		newValue = strings.Replace(newValue, "\n", "\\n", -1)
		newValue = strings.Replace(newValue, "\"", "\\\"", -1)
		newValue = "\"" + newValue + "\""
		return newValue, nil
	default:
		return origin, nil
	}
}

func (b *stringWrap) FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	if value, ok := origin.(string); ok {
		return value, nil
	}
	return "", errors.New("origin content not a string type")
}
