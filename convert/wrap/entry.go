package wrap

import (
	"errors"
	"fmt"
	"table-export/config"
	"table-export/convert/api"
	"table-export/meta"
)

var valueWrapMap map[meta.EFieldType]api.IValueWarp

func init() {
	valueWrapMap = make(map[meta.EFieldType]api.IValueWarp)
	valueWrapMap[meta.FieldType_Int] = &intWrap{}
	valueWrapMap[meta.FieldType_UInt] = &uintWrap{}
	valueWrapMap[meta.FieldType_Long] = &longWrap{}
	valueWrapMap[meta.FieldType_ULong] = &ulongWrap{}
	valueWrapMap[meta.FieldType_Bool] = &boolWrap{}
	valueWrapMap[meta.FiledType_Float] = &floatWrap{}
	valueWrapMap[meta.FiledType_Double] = &doubleWrap{}
	valueWrapMap[meta.FieldType_String] = &stringWrap{}
	valueWrapMap[meta.FieldType_Slice] = &sliceWrap{}
	valueWrapMap[meta.FieldType_Map] = &mapWrap{}
}

func GetOutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return nil, errors.New(fmt.Sprintf("GetOutputValue no support field type[%v]", filedType.Type))
	}
	result, err := wrap.OutputValue(exportType, filedType, origin)
	return result, err
}

func GetOutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetOutputStringValue no support field type[%v]", filedType.Type))
	}
	result, err := wrap.OutputStringValue(exportType, filedType, origin)
	return result, err
}

func GetOutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return "", errors.New(fmt.Sprintf("GetOutputDefTypeValue no support field type[%v]", filedType.Type))
	}
	result, err := wrap.OutputDefTypeValue(exportType, filedType, collectionReadonly)
	return result, err
}

func GetDataVisitorValue(visitor api.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	wrap, ok := valueWrapMap[filedType.Type]
	if !ok {
		return errors.New(fmt.Sprintf("DataVisitorValue no support field type[%v]", filedType.Type))
	}
	err := wrap.DataVisitorValue(visitor, filedType, origin)
	return err
}

func GetCodePrintValue(print api.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	wrap, ok := valueWrapMap[fieldType.Type]
	if !ok {
		return ""
	}
	result := wrap.CodePrintValue(print, fieldType, fieldName, reader, depth)
	return result
}
