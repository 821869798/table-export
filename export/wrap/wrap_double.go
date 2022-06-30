package wrap

import (
	"errors"
	"strconv"
	"table-export/define"
	"table-export/meta"
)

type doubleWrap struct{}

func (b *doubleWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseFloat(origin, 64)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return float64(0), nil
		}
		value, err := strconv.ParseFloat(origin, 64)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *doubleWrap) FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	if value, ok := origin.(float64); ok {
		result := strconv.FormatFloat(value, 'f', -1, 64)
		return result, nil
	}
	return "", errors.New("origin content not a float type")
}
