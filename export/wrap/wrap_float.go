package wrap

import (
	"strconv"
	"table-export/export/common"
	"table-export/meta"
)

type floatWrap struct{}

func (b *floatWrap) OutputValue(exportType common.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case common.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseFloat(origin, 64)
		if err != nil {
			return nil, err
		}
		return origin, nil
	}
	return nil, nil
}
