package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/field_type"
	"strconv"
	"strings"
)

type enumWrap struct{}

func convertEnumToInt(fieldType *field_type.TableFieldType, origin string) (int32, error) {
	enumDefine, ok := env.GetEnumDefine(fieldType.Name)
	if !ok {
		return 0, errors.New(fmt.Sprintf("no support enum type[%v]", fieldType.Name))
	}
	if enumDefine.Parse4String {
		enumValue, ok := enumDefine.ValueStrMapping[strings.TrimSpace(origin)]
		if !ok {
			return 0, errors.New(fmt.Sprintf("no support enum type[%v] value[%v]", fieldType.Name, origin))
		}
		return enumValue.Index, nil
	}
	if origin == "" {
		return int32(0), nil
	}
	value, err := strconv.Atoi(origin)
	if err != nil {
		return 0, err
	}
	result := int32(value)
	_, ok = enumDefine.ValueMapping[result]
	if !ok {
		return 0, errors.New(fmt.Sprintf("no support enum type[%v] value[%v]", fieldType.Name, origin))
	}
	return result, nil
}

func (e *enumWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	return convertEnumToInt(fieldType, origin)
}

func (e *enumWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		value, err := convertEnumToInt(fieldType, origin)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(int(value)), nil
	}
}

func (e *enumWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return env.GetMetaRuleUnitPlus().GetEnumDefinePrefix() + fieldType.Name, nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (e *enumWrap) FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error) {
	switch value := origin.(type) {
	case int32:
		return value, nil
	case int64:
		return int32(value), nil
	case int:
		return int32(value), nil
	case float64:
		return int32(value), nil
	default:
		return nil, errors.New(fmt.Sprintf("[FormatValueInterface|enum] no support type[%T]", origin))
	}
}

func (e *enumWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	value, err := convertEnumToInt(fieldType, origin)
	if err != nil {
		return err
	}
	visitor.AcceptInt(value)
	return nil
}

func (e *enumWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case int32:
		visitor.AcceptInt(value)
		return nil
	case int64:
		visitor.AcceptInt(int32(value))
		return nil
	case int:
		visitor.AcceptInt(int32(value))
		return nil
	case float64:
		visitor.AcceptInt(int32(value))
		return nil
	case string:
		return e.DataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|enum] no support type[%T]", origin))
	}
}

func (e *enumWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptEnum(fieldType, fieldName, reader, depth)
}
