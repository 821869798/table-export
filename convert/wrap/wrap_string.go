package wrap

import (
	"errors"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strings"
)

type stringWrap struct{}

func (b *stringWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
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

func (b *stringWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	if value, ok := origin.(string); ok {
		return value, nil
	}
	return "", errors.New("origin content not a string type")
}

func (b *stringWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		return "string", nil
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *stringWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	visitor.AcceptString(origin)
	return nil
}

func (b *stringWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptString(fieldType, fieldName, reader, depth)
}
