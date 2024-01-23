package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"strconv"
)

type doubleWrap struct{}

func (b *doubleWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return float64(0), nil
	}
	value, err := strconv.ParseFloat(origin, 64)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *doubleWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseFloat(origin, 64)
		if err != nil {
			return "", err
		}
		return origin, nil
	}
}

func (b *doubleWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "double", nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *doubleWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptDouble(0)
		return nil
	}
	value, err := strconv.ParseFloat(origin, 64)
	if err != nil {
		return err
	}
	visitor.AcceptDouble(value)
	return nil
}

func (b *doubleWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case float64:
		visitor.AcceptDouble(value)
		return nil
	case float32:
		visitor.AcceptDouble(float64(value))
		return nil
	case int64:
		visitor.AcceptDouble(float64(value))
		return nil
	case string:
		return b.DataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|double] no support type[%T]", origin))
	}
}

func (b *doubleWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptDouble(fieldType, fieldName, reader, depth)
}
