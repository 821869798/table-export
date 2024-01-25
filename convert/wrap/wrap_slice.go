package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"strings"
)

type sliceWrap struct{}

func (b *sliceWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	origin = strings.TrimSpace(origin)
	strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
	result := make([]interface{}, 0, len(strSlice))
	if origin != "" {
		for _, v := range strSlice {
			content, err := GetOutputValue(exportType, fieldType.Value, v)
			if err != nil {
				return nil, err
			}
			result = append(result, content)
		}
	}
	return result, nil
}

func (b *sliceWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	case config.ExportType_Lua:
		origin = strings.TrimSpace(origin)
		strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
		result := "{"
		if origin != "" {
			for _, v := range strSlice {
				content, err := GetOutputStringValue(exportType, fieldType.Value, v)
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

func (b *sliceWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		valueDef, err := GetOutputDefTypeValue(exportType, fieldType.Value, collectionReadonly)
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
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *sliceWrap) FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error) {
	switch origin.(type) {
	case []interface{}:
		return origin, nil
	case []string:
		return origin, nil
	default:
		return nil, errors.New(fmt.Sprintf("[FormatValueInterface|slice] no support type[%T]", origin))
	}
}

func (b *sliceWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	origin = strings.TrimSpace(origin)
	if origin == "" {
		visitor.AcceptStringArray(EmptyStringArray, fieldType.Value)
		return nil
	}
	strSlice := strings.Split(origin, config.GlobalConfig.Table.ArraySplit)
	visitor.AcceptStringArray(strSlice, fieldType.Value)
	return nil
}

func (b *sliceWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case []interface{}:
		visitor.AcceptArray(value, fieldType.Value)
		return nil
	case []string:
		visitor.AcceptStringArray(value, fieldType.Value)
		return nil
	case string:
		return b.DataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|slice] no support type[%T]", origin))
	}
}

func (b *sliceWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptArray(fieldType, fieldName, reader, depth)
}
