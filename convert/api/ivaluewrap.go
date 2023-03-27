package api

import (
	"table-export/config"
	"table-export/meta"
)

type IValueWarp interface {
	OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error)
	OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error)
	OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error)
	DataVisitorValue(visitor IDataVisitor, filedType *meta.TableFieldType, origin string) error
	CodePrintValue(print ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
}
