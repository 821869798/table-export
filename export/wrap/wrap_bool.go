package wrap

import (
	"errors"
	"strconv"
	"table-export/define"
	"table-export/meta"
)

type boolWrap struct{}

func (b *boolWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		if origin == "" {
			return strconv.FormatBool(false), nil
		}
		value, err := strconv.ParseBool(origin)
		if err != nil {
			return nil, err
		}
		return strconv.FormatBool(value), nil
	default:
		if origin == "" {
			return false, nil
		}
		value, err := strconv.ParseBool(origin)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *boolWrap) FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	if value, ok := origin.(bool); ok {
		result := strconv.FormatBool(value)
		return result, nil
	}
	return "", errors.New("origin content not a bool type")
}
