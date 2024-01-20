package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
)

var valueWrapMap map[meta.EFieldType]IValueWrap

func init() {
	valueWrapMap = make(map[meta.EFieldType]IValueWrap)
	valueWrapMap[meta.EFieldType_Int] = &intWrap{}
	valueWrapMap[meta.EFieldType_UInt] = &uintWrap{}
	valueWrapMap[meta.EFieldType_Long] = &longWrap{}
	valueWrapMap[meta.EFieldType_ULong] = &ulongWrap{}
	valueWrapMap[meta.EFieldType_Bool] = &boolWrap{}
	valueWrapMap[meta.EFieldType_Float] = &floatWrap{}
	valueWrapMap[meta.EFieldType_Double] = &doubleWrap{}
	valueWrapMap[meta.EFieldType_String] = &stringWrap{}
	valueWrapMap[meta.EFieldType_Slice] = &sliceWrap{}
	valueWrapMap[meta.EFieldType_Map] = &mapWrap{}
	valueWrapMap[meta.EFieldType_Enum] = &enumWrap{}
	valueWrapMap[meta.EFieldType_Class] = &classWrap{}
}

// GetOutputValue 获取真实的值，使用interface{}包装
func GetOutputValue(exportType config.ExportType, fieldType *meta.TableFieldType, origin string) (interface{}, error) {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return nil, errors.New(fmt.Sprintf("GetOutputValue no support field type[%v]", fieldType.Type))
	}
	result, err := wrap.OutputValue(exportType, fieldType, origin)
	return result, err
}

// GetOutputStringValue 获取该导出类型下String的值
func GetOutputStringValue(exportType config.ExportType, fieldType *meta.TableFieldType, origin string) (string, error) {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetOutputStringValue no support field type[%v]", fieldType.Type))
	}
	result, err := wrap.OutputStringValue(exportType, fieldType, origin)
	return result, err
}

// GetOutputDefTypeValue 获取这种类型的导出的定义符号·
func GetOutputDefTypeValue(exportType config.ExportType, fieldType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetOutputDefTypeValue no support field type[%v]", fieldType.Type))
	}
	result, err := wrap.OutputDefTypeValue(exportType, fieldType, collectionReadonly)
	return result, err
}

func RunDataVisitorString(visitor apiconvert.IDataVisitor, fieldType *meta.TableFieldType, origin string) error {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return errors.New(fmt.Sprintf("DataVisitorString no support field type[%v]", fieldType.Type))
	}
	err := wrap.DataVisitorString(visitor, fieldType, origin)
	return err
}

func RunDataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *meta.TableFieldType, origin interface{}) error {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return errors.New(fmt.Sprintf("DataVisitorString no support field type[%v]", fieldType.Type))
	}
	err := wrap.DataVisitorValue(visitor, fieldType, origin)
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
