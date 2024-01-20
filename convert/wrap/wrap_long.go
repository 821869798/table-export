package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strconv"
)

type longWrap struct{}

func (b *longWrap) OutputValue(exportType config.ExportType, fieldType *meta.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return int64(0), nil
	}
	value, err := strconv.ParseInt(origin, 10, 64)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *longWrap) OutputStringValue(exportType config.ExportType, fieldType *meta.TableFieldType, origin string) (string, error) {
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

func (b *longWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "long", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *longWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *meta.TableFieldType, origin string) error {
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

func (b *longWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *meta.TableFieldType, origin interface{}) error {
	value, ok := origin.(int64)
	if ok {
		visitor.AcceptLong(value)
		return nil
	}
	stringValue, ok := origin.(string)
	if ok {
		return b.DataVisitorString(visitor, fieldType, stringValue)
	}
	return errors.New(fmt.Sprintf("[DataVisitorValue|long] no support type[%T]", origin))
}

func (b *longWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptLong(fieldType, fieldName, reader, depth)
}
