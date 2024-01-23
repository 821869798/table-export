package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"strconv"
)

type ulongWrap struct{}

func (b *ulongWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return uint64(0), nil
	}
	value, err := strconv.ParseUint(origin, 10, 64)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *ulongWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseUint(origin, 10, 64)
		if err != nil {
			return "", err
		}
		return origin, nil
	}
}

func (b *ulongWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "ulong", nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *ulongWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptULong(0)
		return nil
	}
	value, err := strconv.ParseUint(origin, 10, 64)
	if err != nil {
		return err
	}
	visitor.AcceptULong(value)
	return nil
}

func (b *ulongWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case uint64:
		visitor.AcceptULong(value)
		return nil
	case uint32:
		visitor.AcceptULong(uint64(value))
		return nil
	case int:
		visitor.AcceptULong(uint64(value))
		return nil
	case float64:
		visitor.AcceptULong(uint64(value))
		return nil
	case int64:
		visitor.AcceptULong(uint64(value))
		return nil
	case int32:
		visitor.AcceptULong(uint64(value))
		return nil
	case string:
		return b.DataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|ulong] no support type[%T]", origin))
	}
}

func (b *ulongWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptULong(fieldType, fieldName, reader, depth)
}
