package wrap

import (
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
	}
	return nil, nil
}
