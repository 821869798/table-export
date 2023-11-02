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

type floatWrap struct{}

func (b *floatWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
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

func (b *floatWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (string, error) {
	switch exportType {
	default:
		if origin == "" {
			return "0", nil
		}
		_, err := strconv.ParseFloat(origin, 32)
		if err != nil {
			return "", err
		}
		return origin, nil
	}
}

func (b *floatWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "float", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *floatWrap) DataVisitorString(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
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

func (b *floatWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case float32:
		visitor.AcceptFloat(value)
		return nil
	case float64:
		visitor.AcceptFloat(float32(value))
		return nil
	case string:
		return b.DataVisitorString(visitor, filedType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|float] no support type[%T]", origin))
	}
}

func (b *floatWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptFloat(fieldType, fieldName, reader, depth)
}
