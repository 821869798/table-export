package wrap

import (
	"errors"
	"math"
	"strconv"
	"table-export/config"
	"table-export/convert/api"
	"table-export/meta"
)

type floatWrap struct{}

func (b *floatWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	switch exportType {
	case config.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseFloat(origin, 32)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return float32(0), nil
		}
		value, err := strconv.ParseFloat(origin, 32)
		if err != nil {
			return nil, err
		}
		if value > math.MaxFloat32 {
			return nil, errors.New("float value can't greater than max float32")
		}
		return float32(value), nil
	}
}

func (b *floatWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	if value, ok := origin.(float32); ok {
		result := strconv.FormatFloat(float64(value), 'f', -1, 64)
		return result, nil
	}
	return "", errors.New("origin content not a float type")
}

func (b *floatWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "float", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *floatWrap) DataVisitorValue(visitor api.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptFloat(0)
		return nil
	}
	value, err := strconv.ParseFloat(origin, 32)
	if err != nil {
		return err
	}
	visitor.AcceptFloat(float32(value))
	return nil
}

func (b *floatWrap) CodePrintValue(print api.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptFloat(fieldType, fieldName, reader, depth)
}
