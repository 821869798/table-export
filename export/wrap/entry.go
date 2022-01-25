package wrap

import (
	"errors"
	"fmt"
	"table-export/define"
	"table-export/meta"
)

type ValueWarp interface {
	OutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error)
	FormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error)
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

func GetOutputValue(exportType define.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return nil, errors.New(fmt.Sprintf("GetOutputValue no support field type[%v]", filedType.Type))
	}
	result, err := wrap.OutputValue(exportType, filedType, origin)
	return result, err
}

func GetFormatValue(exportType define.ExportType, filedType *meta.TableFiledType, origin interface{}) (string, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetFormatValue no support field type[%v]", filedType.Type))
	}
	result, err := wrap.FormatValue(exportType, filedType, origin)
	return result, err
}
