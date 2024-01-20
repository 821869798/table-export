package wrap

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
)

type IValueWrap interface {
	OutputValue(exportType config.ExportType, fieldType *meta.TableFieldType, origin string) (interface{}, error)
	OutputStringValue(exportType config.ExportType, fieldType *meta.TableFieldType, origin string) (string, error)
	OutputDefTypeValue(exportType config.ExportType, fieldType *meta.TableFieldType, collectionReadonly bool) (string, error)
	DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *meta.TableFieldType, origin string) error
	DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *meta.TableFieldType, origin interface{}) error
	CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
}
