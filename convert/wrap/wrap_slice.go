package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strings"
)

type sliceWrap struct{}

func (b *sliceWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
	result := make([]interface{}, 0, len(strSlice))
	if origin != "" {
		for _, v := range strSlice {
			content, err := GetOutputValue(exportType, filedType.Value, v)
			if err != nil {
				return nil, err
			}
			result = append(result, content)
		}
	}
	return result, nil
}

func (b *sliceWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (string, error) {
	switch exportType {
	case config.ExportType_Lua:
		strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
		result := "{"
		if origin != "" {
			for _, v := range strSlice {
				content, err := GetOutputStringValue(exportType, filedType.Value, v)
				if err != nil {
					return "", err
				}
				result += content + ","
			}
		}
		result += "}"
		return result, nil
	default:
		return "", errors.New(fmt.Sprintf("OutputStringValue slice no support export type[%v]", exportType))
	}
}

func (b *sliceWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		valueDef, err := GetOutputDefTypeValue(exportType, filedType.Value, collectionReadonly)
		if err != nil {
			return "", err
		}
		var result string
		if collectionReadonly {
			result = fmt.Sprintf("IReadOnlyList<%s>", valueDef)
		} else {
			result = fmt.Sprintf("%s[]", valueDef)
		}

		return result, nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *sliceWrap) DataVisitorString(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptStringArray(EmptyStringArray, filedType.Value)
		return nil
	}
	strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
	visitor.AcceptStringArray(strSlice, filedType.Value)
	return nil
}

func (b *sliceWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case []interface{}:
		visitor.AcceptArray(value, filedType.Value)
		return nil
	case []string:
		visitor.AcceptStringArray(value, filedType.Value)
		return nil
	case string:
		return b.DataVisitorString(visitor, filedType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|int] no support type[%T]", origin))
	}
}

func (b *sliceWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptArray(fieldType, fieldName, reader, depth)
}
