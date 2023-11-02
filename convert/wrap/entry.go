package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
)

var valueWrapMap map[meta.EFieldType]IValueWarp

func init() {
	valueWrapMap = make(map[meta.EFieldType]IValueWarp)
	valueWrapMap[meta.EFieldType_Int] = &intWrap{}
	valueWrapMap[meta.EFieldType_UInt] = &uintWrap{}
	valueWrapMap[meta.EFieldType_Long] = &longWrap{}
	valueWrapMap[meta.EFieldType_ULong] = &ulongWrap{}
	valueWrapMap[meta.EFieldType_Bool] = &boolWrap{}
	valueWrapMap[meta.EFiledType_Float] = &floatWrap{}
	valueWrapMap[meta.EFiledType_Double] = &doubleWrap{}
	valueWrapMap[meta.EFieldType_String] = &stringWrap{}
	valueWrapMap[meta.EFieldType_Slice] = &sliceWrap{}
	valueWrapMap[meta.EFieldType_Map] = &mapWrap{}
}

// GetOutputValue 获取真实的值，使用interface{}包装
func GetOutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return nil, errors.New(fmt.Sprintf("GetOutputValue no support field type[%v]", filedType.Type))
	}
	result, err := wrap.OutputValue(exportType, filedType, origin)
	return result, err
}

// GetOutputStringValue 获取该导出类型下String的值
func GetOutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (string, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetOutputStringValue no support field type[%v]", filedType.Type))
	}
	result, err := wrap.OutputStringValue(exportType, filedType, origin)
	return result, err
}

// GetOutputDefTypeValue 获取这种类型的导出的定义符号·
func GetOutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetOutputDefTypeValue no support field type[%v]", filedType.Type))
	}
	result, err := wrap.OutputDefTypeValue(exportType, filedType, collectionReadonly)
	return result, err
}

func RunDataVisitorString(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return errors.New(fmt.Sprintf("DataVisitorString no support field type[%v]", filedType.Type))
	}
	err := wrap.DataVisitorString(visitor, filedType, origin)
	return err
}

func RunDataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin interface{}) error {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return errors.New(fmt.Sprintf("DataVisitorString no support field type[%v]", filedType.Type))
	}
	err := wrap.DataVisitorValue(visitor, filedType, origin)
	return err
}

func GetCodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return ""
	}
	result := wrap.CodePrintValue(print, fieldType, fieldName, reader, depth)
	return result
}
