package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"math"
	"strconv"
)

type uintWrap struct{}

func (b *uintWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
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

func (b *uintWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseUint(origin, 10, 32)
		if err != nil {
			return "", err
		}
		return origin, nil
	}
}

func (b *uintWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "uint", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *uintWrap) DataVisitorString(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
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

func (b *uintWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case uint32:
		visitor.AcceptUInt(value)
		return nil
	case uint64:
		visitor.AcceptUInt(uint32(value))
		return nil
	case string:
		return b.DataVisitorString(visitor, filedType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|int] no support type[%T]", origin))
	}
}

func (b *uintWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptUInt(fieldType, fieldName, reader, depth)
}
