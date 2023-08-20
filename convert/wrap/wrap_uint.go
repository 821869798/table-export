package wrap

import (
	"errors"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"math"
	"strconv"
)

type uintWrap struct{}

func (b *uintWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
	switch exportType {
	case config.ExportType_Lua:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseUint(origin, 10, 32)
		if err != nil {
			return nil, err
		}
		return origin, nil
	default:
		if origin == "" {
			return uint32(0), nil
		}
		value, err := strconv.ParseUint(origin, 10, 32)
		if err != nil {
			return nil, err
		}
		if value > math.MaxUint32 {
			return nil, errors.New("uint value can't greater than max uint32")
		}
		return uint32(value), nil
	}
}

func (b *uintWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	if value, ok := origin.(uint32); ok {
		result := strconv.FormatUint(uint64(value), 10)
		return result, nil
	}
	return "", errors.New("origin content not a uint type")
}

func (b *uintWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "uint", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *uintWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptUInt(0)
		return nil
	}
	value, err := strconv.ParseUint(origin, 10, 32)
	if err != nil {
		return err
	}
	visitor.AcceptUInt(uint32(value))
	return nil
}

func (b *uintWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptUInt(fieldType, fieldName, reader, depth)
}
