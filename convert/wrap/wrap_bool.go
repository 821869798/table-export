package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"strconv"
)

type boolWrap struct{}

func (b *boolWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return false, nil
	}
	value, err := strconv.ParseBool(origin)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *boolWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		if origin == "" {
			return strconv.FormatBool(false), nil
		}
		value, err := strconv.ParseBool(origin)
		if err != nil {
			return "", err
		}
		return strconv.FormatBool(value), nil
	}
}

func (b *boolWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "bool", nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *boolWrap) FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error) {
	_, ok := origin.(bool)
	if ok {
		return origin, nil
	}
	return nil, errors.New(fmt.Sprintf("[FormatValueInterface|bool] no support type[%T]", origin))
}

func (b *boolWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptBool(false)
		return nil
	}
	value, err := strconv.ParseBool(origin)
	if err != nil {
		return err
	}
	visitor.AcceptBool(value)
	return nil
}

func (b *boolWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	value, ok := origin.(bool)
	if ok {
		visitor.AcceptBool(value)
		return nil
	}
	stringValue, ok := origin.(string)
	if ok {
		return b.DataVisitorString(visitor, fieldType, stringValue)
	}
	return errors.New(fmt.Sprintf("[DataVisitorValue|bool] no support type[%T]", origin))
}

func (b *boolWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptBool(fieldType, fieldName, reader, depth)
}
