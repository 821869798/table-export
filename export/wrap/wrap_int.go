package wrap

import (
	"errors"
	"strconv"
	"table-export/define"
	"table-export/meta"
)

type intWrap struct{}

func (b *intWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.Atoi(origin)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return 0, nil
		}
		value, err := strconv.Atoi(origin)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *intWrap) FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	if value, ok := origin.(int); ok {
		result := strconv.Itoa(value)
		return result, nil
	}
	return "", errors.New("origin content not a int type")
}
