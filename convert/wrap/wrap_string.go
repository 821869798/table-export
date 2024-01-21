package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"strings"
)

type stringWrap struct{}

func (b *stringWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	return origin, nil
}

func (b *stringWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	switch exportType {
	case config.ExportType_Lua:
		newValue := strings.Replace(origin, "\\", "\\\\", -1)
		newValue = strings.Replace(newValue, "\n", "\\n", -1)
		newValue = strings.Replace(newValue, "\"", "\\\"", -1)
		newValue = "\"" + newValue + "\""
		return newValue, nil
	default:
		return origin, nil
	}
}

func (b *stringWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "string", nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *stringWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	visitor.AcceptString(origin)
	return nil
}

func (b *stringWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	stringValue, ok := origin.(string)
	if ok {
		visitor.AcceptString(stringValue)
		return nil
	}
	return errors.New(fmt.Sprintf("[DataVisitorValue|string] no support type[%T]", origin))
}

func (b *stringWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptString(fieldType, fieldName, reader, depth)
}
