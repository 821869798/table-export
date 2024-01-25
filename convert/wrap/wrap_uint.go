package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"math"
	"strconv"
)

type uintWrap struct{}

func (b *uintWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return uint32(0), nil
	}
	value, err := strconv.ParseUint(origin, 10, 32)
	if err != nil {
		return nil, err
	}
	if value > math.MaxUint32 {
		return nil, errors.New("uint value can't greater than max uint32")
	}
	return uint32(value), nil
}

func (b *uintWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseUint(origin, 10, 32)
		if err != nil {
			return "", err
		}
		return origin, nil
	}
}

func (b *uintWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "uint", nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *uintWrap) FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error) {
	switch value := origin.(type) {
	case uint32:
		return value, nil
	case uint64:
		return uint32(value), nil
	case int:
		return uint32(value), nil
	case float64:
		return uint32(value), nil
	case int64:
		return uint32(value), nil
	case int32:
		return uint32(value), nil
	default:
		return nil, errors.New(fmt.Sprintf("[FormatValueInterface|uint] no support type[%T]", origin))
	}
}

func (b *uintWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptUInt(0)
		return nil
	}
	value, err := strconv.ParseUint(origin, 10, 32)
	if err != nil {
		return err
	}
	visitor.AcceptUInt(uint32(value))
	return nil
}

func (b *uintWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case uint32:
		visitor.AcceptUInt(value)
		return nil
	case uint64:
		visitor.AcceptUInt(uint32(value))
		return nil
	case int:
		visitor.AcceptUInt(uint32(value))
		return nil
	case float64:
		visitor.AcceptUInt(uint32(value))
		return nil
	case int64:
		visitor.AcceptUInt(uint32(value))
		return nil
	case int32:
		visitor.AcceptUInt(uint32(value))
		return nil
	case string:
		return b.DataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|uint] no support type[%T]", origin))
	}
}

func (b *uintWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptUInt(fieldType, fieldName, reader, depth)
}
