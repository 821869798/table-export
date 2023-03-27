package wrap

import (
	"errors"
	"strconv"
	"table-export/config"
	"table-export/convert/api"
	"table-export/meta"
)

type longWrap struct{}

func (b *longWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	switch exportType {
	case config.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseInt(origin, 10, 64)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return int64(0), nil
		}
		value, err := strconv.ParseInt(origin, 10, 64)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *longWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	if value, ok := origin.(int64); ok {
		result := strconv.FormatInt(value, 10)
		return result, nil
	}
	return "", errors.New("origin content not a int type")
}

func (b *longWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "long", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *longWrap) DataVisitorValue(visitor api.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
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

func (b *longWrap) CodePrintValue(print api.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptLong(fieldType, fieldName, reader, depth)
}
