package wrap

import (
	"errors"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strconv"
)

type boolWrap struct{}

func (b *boolWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	switch exportType {
	case config.ExportType_Lua:
		if origin == "" {
			return strconv.FormatBool(false), nil
		}
		value, err := strconv.ParseBool(origin)
		if err != nil {
			return nil, err
		}
		return strconv.FormatBool(value), nil
	default:
		if origin == "" {
			return false, nil
		}
		value, err := strconv.ParseBool(origin)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *boolWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	if value, ok := origin.(bool); ok {
		result := strconv.FormatBool(value)
		return result, nil
	}
	return "", errors.New("origin content not a bool type")
}

func (b *boolWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "bool", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *boolWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
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

func (b *boolWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptBool(fieldType, fieldName, reader, depth)
}
