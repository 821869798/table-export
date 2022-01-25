package wrap

import (
	"errors"
	"fmt"
	"table-export/export/common"
	"table-export/meta"
)

type ValueWarp interface {
	OutputValue(exportType common.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error)
}

var valueWrapMap map[meta.FieldType]ValueWarp

func init() {
	valueWrapMap = make(map[meta.FieldType]ValueWarp)
	valueWrapMap[meta.FieldType_Int] = &intWrap{}
	valueWrapMap[meta.FieldType_Bool] = &boolWrap{}
	valueWrapMap[meta.FiledType_Float] = &floatWrap{}
	valueWrapMap[meta.FieldType_String] = &stringWrap{}
	valueWrapMap[meta.FieldType_Slice] = &sliceWrap{}
	valueWrapMap[meta.FieldType_Map] = &mapWrap{}
}

func GetOutputValue(exportType common.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no support field type[%v]", filedType.Type))
	}
	result, err := wrap.OutputValue(exportType, filedType, origin)
	return result, err
}
