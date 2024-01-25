package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/field_type"
)

type classWrap struct{}

func (c *classWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	//TODO implement me
	panic("classWrap implement me")
}

func (c *classWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	//TODO implement me
	panic("classWrap implement me")
}

func (c *classWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	default:
		return env.GetMetaRuleUnitPlus().GetClassDefinePrefix() + fieldType.Name, nil
	}
}

func (c *classWrap) FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error) {
	if origin == nil {
		return nil, nil
	}
	switch origin.(type) {
	case map[string]interface{}:
		return origin, nil
	case map[string]string:
		return origin, nil
	default:
		return nil, errors.New(fmt.Sprintf("[FormatValueInterface|class] no support type[%T]", origin))
	}
}

func (c *classWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	//TODO implement me
	panic("implement me")
}

func (c *classWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	if origin == nil {
		visitor.AcceptClassNull(fieldType.Class)
		return nil
	}
	switch value := origin.(type) {
	case map[string]interface{}:
		visitor.AcceptClass(value, fieldType.Class)
		return nil
	case map[string]string:
		visitor.AcceptClassString(value, fieldType.Class)
		return nil
	case string:
		return RunDataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|class] no support type[%T]", origin))
	}
}

func (c *classWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptClass(fieldType, fieldName, reader, depth)
}
