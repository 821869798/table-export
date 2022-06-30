package wrap

import (
	"errors"
	"strconv"
	"table-export/define"
	"table-export/meta"
)

type longWrap struct{}

func (b *longWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseInt(origin, 10, 64)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return int64(0), nil
		}
		value, err := strconv.ParseInt(origin, 10, 64)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *longWrap) FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	if value, ok := origin.(int64); ok {
		result := strconv.FormatInt(value, 10)
		return result, nil
	}
	return "", errors.New("origin content not a int type")
}
