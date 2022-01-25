package wrap

import (
	"strconv"
	"table-export/export/common"
	"table-export/meta"
)

type intWrap struct{}

func (b *intWrap) OutputValue(exportType common.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case common.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.Atoi(origin)
		if err != nil {
			return nil, err
		}
		return origin, nil
	}
	return nil, nil
}
