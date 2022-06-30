package wrap

import (
	"errors"
	"strconv"
	"table-export/define"
	"table-export/meta"
)

type ulongWrap struct{}

func (b *ulongWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseUint(origin, 10, 64)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return uint64(0), nil
		}
		value, err := strconv.ParseUint(origin, 10, 64)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *ulongWrap) FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	if value, ok := origin.(uint64); ok {
		result := strconv.FormatUint(value, 10)
		return result, nil
	}
	return "", errors.New("origin content not a uint type")
}
