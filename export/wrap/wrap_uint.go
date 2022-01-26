package wrap

import (
	"errors"
	"math"
	"strconv"
	"table-export/define"
	"table-export/meta"
)

type uintWrap struct{}

func (b *uintWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseUint(origin, 10, 32)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return uint32(0), nil
		}
		value, err := strconv.ParseUint(origin, 10, 32)
		if err != nil {
			return nil, err
		}
		if value > math.MaxUint32 {
			return nil, errors.New("uint value can't greater than max uint32")
		}
		return uint32(value), nil
	}
}

func (b *uintWrap) FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	if value, ok := origin.(uint32); ok {
		result := strconv.FormatUint(uint64(value), 10)
		return result, nil
	}
	return "", errors.New("origin content not a uint type")
}
