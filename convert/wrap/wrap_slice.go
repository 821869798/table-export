package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/adapter"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strings"
)

type sliceWrap struct{}

func (b *sliceWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	switch exportType {
	case config.ExportType_Lua:
		strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
		result := "{"
		if origin != "" {
			for _, v := range strSlice {
				content, err := GetOutputValue(exportType, filedType.Value, v)
				if err != nil {
					return nil, err
				}
				valueStr, ok := content.(string)
				if !ok {
					return nil, errors.New("slice get output value error")
				}
				result += valueStr + ","
			}
		}
		result += "}"
		return result, nil
	default:
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
}

func (b *sliceWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	return "", errors.New("slice no support OutputStringValue")
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

func (b *sliceWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptArray(adapter.EmptyArray)
		return nil
	}
	strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
	array := adapter.NewArray(strSlice, filedType.Value)
	visitor.AcceptArray(array)
	return nil
}

func (b *sliceWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptArray(fieldType, fieldName, reader, depth)
}
