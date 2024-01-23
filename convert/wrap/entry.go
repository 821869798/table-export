package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
)

var valueWrapMap map[field_type.EFieldType]IValueWrap

func init() {
	valueWrapMap = make(map[field_type.EFieldType]IValueWrap)
	valueWrapMap[field_type.EFieldType_Int] = &intWrap{}
	valueWrapMap[field_type.EFieldType_UInt] = &uintWrap{}
	valueWrapMap[field_type.EFieldType_Long] = &longWrap{}
	valueWrapMap[field_type.EFieldType_ULong] = &ulongWrap{}
	valueWrapMap[field_type.EFieldType_Bool] = &boolWrap{}
	valueWrapMap[field_type.EFieldType_Float] = &floatWrap{}
	valueWrapMap[field_type.EFieldType_Double] = &doubleWrap{}
	valueWrapMap[field_type.EFieldType_String] = &stringWrap{}
	valueWrapMap[field_type.EFieldType_Slice] = &sliceWrap{}
	valueWrapMap[field_type.EFieldType_Map] = &mapWrap{}
	valueWrapMap[field_type.EFieldType_Enum] = &enumWrap{}
	valueWrapMap[field_type.EFieldType_Class] = &classWrap{}
}

// GetOutputValue 获取真实的值，使用interface{}包装
func GetOutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	if fieldType.ExtFieldType != nil {
		return fieldType.ExtFieldType.ParseOriginData(origin)
	}
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return nil, errors.New(fmt.Sprintf("GetOutputValue no support field_type type[%v]", fieldType.Type))
	}
	result, err := wrap.OutputValue(exportType, fieldType, origin)
	return result, err
}

// GetOutputStringValue 获取该导出类型下String的值
func GetOutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetOutputStringValue no support field_type type[%v]", fieldType.Type))
	}
	result, err := wrap.OutputStringValue(exportType, fieldType, origin)
	return result, err
}

// GetOutputDefTypeValue 获取这种类型的导出的定义符号·
func GetOutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetOutputDefTypeValue no support field_type type[%v]", fieldType.Type))
	}
	result, err := wrap.OutputDefTypeValue(exportType, fieldType, collectionReadonly)
	return result, err
}

func RunDataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	if fieldType.ExtFieldType != nil {
		data, err := fieldType.ExtFieldType.ParseOriginData(origin)
		if err != nil {
			return err
		}
		return RunDataVisitorValue(visitor, fieldType, data)
	}
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return errors.New(fmt.Sprintf("DataVisitorString no support field_type type[%v]", fieldType.Type))
	}
	err := wrap.DataVisitorString(visitor, fieldType, origin)
	return err
}

func RunDataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return errors.New(fmt.Sprintf("DataVisitorString no support field_type type[%v]", fieldType.Type))
	}
	err := wrap.DataVisitorValue(visitor, fieldType, origin)
	return err
}

func GetCodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return ""
	}
	result := wrap.CodePrintValue(print, fieldType, fieldName, reader, depth)
	return result
}
