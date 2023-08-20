package wrap

import (
	"errors"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strconv"
)

type intWrap struct{}

func (b *intWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	switch exportType {
	case config.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.Atoi(origin)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return int32(0), nil
		}
		value, err := strconv.Atoi(origin)
		if err != nil {
			return nil, err
		}
		return int32(value), nil
	}
}

func (b *intWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	if value, ok := origin.(int32); ok {
		result := strconv.Itoa(int(value))
		return result, nil
	}
	return "", errors.New("origin content not a int type")
}

func (b *intWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "int", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *intWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptInt(0)
		return nil
	}
	value, err := strconv.Atoi(origin)
	if err != nil {
		return err
	}
	visitor.AcceptInt(int32(value))
	return nil
}

func (b *intWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptInt(fieldType, fieldName, reader, depth)
}
