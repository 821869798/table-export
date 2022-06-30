package wrap

import (
	"errors"
	"math"
	"strconv"
	"table-export/define"
	"table-export/meta"
)

type floatWrap struct{}

func (b *floatWrap) OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case define.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseFloat(origin, 32)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return float32(0), nil
		}
		value, err := strconv.ParseFloat(origin, 32)
		if err != nil {
			return nil, err
		}
		if value > math.MaxFloat32 {
			return nil, errors.New("float value can't greater than max float32")
		}
		return float32(value), nil
	}
}

func (b *floatWrap) FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	if value, ok := origin.(float32); ok {
		result := strconv.FormatFloat(float64(value), 'f', -1, 64)
		return result, nil
	}
	return "", errors.New("origin content not a float type")
}
