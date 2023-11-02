package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strconv"
)

type doubleWrap struct{}

func (b *doubleWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return float64(0), nil
	}
	value, err := strconv.ParseFloat(origin, 64)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *doubleWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (string, error) {
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

func (b *doubleWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "double", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *doubleWrap) DataVisitorString(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
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

func (b *doubleWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin interface{}) error {
	value, ok := origin.(float64)
	if ok {
		visitor.AcceptDouble(value)
		return nil
	}
	stringValue, ok := origin.(string)
	if ok {
		return b.DataVisitorString(visitor, filedType, stringValue)
	}
	return errors.New(fmt.Sprintf("[DataVisitorValue|double] no support type[%T]", origin))
}

func (b *doubleWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptDouble(fieldType, fieldName, reader, depth)
}
