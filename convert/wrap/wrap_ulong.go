package wrap

import (
	"errors"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strconv"
)

type ulongWrap struct{}

func (b *ulongWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	switch exportType {
	case config.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseUint(origin, 10, 64)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return uint64(0), nil
		}
		value, err := strconv.ParseUint(origin, 10, 64)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *ulongWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	if value, ok := origin.(uint64); ok {
		result := strconv.FormatUint(value, 10)
		return result, nil
	}
	return "", errors.New("origin content not a uint type")
}

func (b *ulongWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "ulong", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *ulongWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
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

func (b *ulongWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptULong(fieldType, fieldName, reader, depth)
}
