package wrap

import (
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
	}
	return nil, nil
}
