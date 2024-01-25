package wrap

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
)

type IValueWrap interface {
	OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error)
	OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error)
	OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error)
	FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error)
	DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error
	DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error
	CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
}
