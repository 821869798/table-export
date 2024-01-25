package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"strconv"
)

type longWrap struct{}

func (b *longWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return int64(0), nil
	}
	value, err := strconv.ParseInt(origin, 10, 64)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *longWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseInt(origin, 10, 64)
		if err != nil {
			return "", err
		}
		return origin, nil
	}
}

func (b *longWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "long", nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *longWrap) FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error) {
	switch value := origin.(type) {
	case int64:
		return value, nil
	case int32:
		return int64(value), nil
	case int:
		return int64(value), nil
	case float64:
		return int64(value), nil
	default:
		return nil, errors.New(fmt.Sprintf("[FormatValueInterface|long] no support type[%T]", origin))
	}
}

func (b *longWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptLong(0)
		return nil
	}
	value, err := strconv.ParseInt(origin, 10, 64)
	if err != nil {
		return err
	}
	visitor.AcceptLong(value)
	return nil
}

func (b *longWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case int64:
		visitor.AcceptLong(value)
		return nil
	case int32:
		visitor.AcceptLong(int64(value))
		return nil
	case int:
		visitor.AcceptLong(int64(value))
		return nil
	case float64:
		visitor.AcceptLong(int64(value))
		return nil
	case string:
		return b.DataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|long] no support type[%T]", origin))
	}
}

func (b *longWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptLong(fieldType, fieldName, reader, depth)
}
