package wrap

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
)

type classWrap struct{}

func (c *classWrap) OutputValue(exportType config.ExportType, fieldType *meta.TableFieldType, origin string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *classWrap) OutputStringValue(exportType config.ExportType, fieldType *meta.TableFieldType, origin string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *classWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *classWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *meta.TableFieldType, origin string) error {
	//TODO implement me
	panic("implement me")
}

func (c *classWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *meta.TableFieldType, origin interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *classWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	//TODO implement me
	panic("implement me")
}
