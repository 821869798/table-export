package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strconv"
)

type ulongWrap struct{}

func (b *ulongWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	if origin == "" {
		return uint64(0), nil
	}
	value, err := strconv.ParseUint(origin, 10, 64)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *ulongWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (string, error) {
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

func (b *ulongWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "ulong", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *ulongWrap) DataVisitorString(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
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

func (b *ulongWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin interface{}) error {
	value, ok := origin.(uint64)
	if ok {
		visitor.AcceptULong(value)
		return nil
	}
	stringValue, ok := origin.(string)
	if ok {
		return b.DataVisitorString(visitor, filedType, stringValue)
	}
	return errors.New(fmt.Sprintf("[DataVisitorValue|long] no support type[%T]", origin))
}

func (b *ulongWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptULong(fieldType, fieldName, reader, depth)
}
