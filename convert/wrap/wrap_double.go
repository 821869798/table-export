package wrap

import (
	"errors"
	"strconv"
	"table-export/config"
	"table-export/convert/api"
	"table-export/meta"
)

type doubleWrap struct{}

func (b *doubleWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	switch exportType {
	case config.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseFloat(origin, 64)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return float64(0), nil
		}
		value, err := strconv.ParseFloat(origin, 64)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}

func (b *doubleWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	if value, ok := origin.(float64); ok {
		result := strconv.FormatFloat(value, 'f', -1, 64)
		return result, nil
	}
	return "", errors.New("origin content not a float type")
}

func (b *doubleWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "double", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *doubleWrap) DataVisitorValue(visitor api.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
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

func (b *doubleWrap) CodePrintValue(print api.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptDouble(fieldType, fieldName, reader, depth)
}
