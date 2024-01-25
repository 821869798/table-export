package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"strconv"
)

type intWrap struct{}

func (b *intWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return int32(0), nil
	}
	value, err := strconv.Atoi(origin)
	if err != nil {
		return nil, err
	}
	return int32(value), nil
}

func (b *intWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.Atoi(origin)
		if err != nil {
			return "", err
		}
		return origin, nil
	}
}

func (b *intWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "int", nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *intWrap) FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error) {
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
		return nil, errors.New(fmt.Sprintf("[FormatValueInterface|int] no support type[%T]", origin))
	}
}

func (b *intWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptInt(0)
		return nil
	}
	value, err := strconv.Atoi(origin)
	if err != nil {
		return err
	}
	visitor.AcceptInt(int32(value))
	return nil
}

func (b *intWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
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
		return b.DataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|int] no support type[%T]", origin))
	}
}

func (b *intWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptInt(fieldType, fieldName, reader, depth)
}
