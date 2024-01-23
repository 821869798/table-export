package table_engine

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/field_type"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	_ "github.com/dop251/goja_nodejs/util"
)

type TableEngine struct {
	runtime *goja.Runtime
	//util        *goja.Object
}

func (c *TableEngine) NewTableFieldClass(name string) *field_type.TableFieldClass {
	return field_type.NewTableFieldClass(name)
}

func (c *TableEngine) NewTableFieldType(fieldType int) *field_type.TableFieldType {
	return field_type.NewTableFieldType(field_type.EFieldType(fieldType))
}

func (c *TableEngine) NewTableFieldEnumType(name string) *field_type.TableFieldType {
	return field_type.NewTableFieldEnumType(name)
}

func (c *TableEngine) NewTableFieldArrayType(value *field_type.TableFieldType) *field_type.TableFieldType {
	return field_type.NewTableFieldArrayType(value)
}

func (c *TableEngine) NewTableFieldMapType(key *field_type.TableFieldType, value *field_type.TableFieldType) *field_type.TableFieldType {
	return field_type.NewTableFieldMapType(key, value)
}

func (c *TableEngine) NewTableFieldClassType(class *field_type.TableFieldClass) *field_type.TableFieldType {
	return field_type.NewTableFieldClassType(class)
}

//func (c *TableEngine) log(p func(string)) func(goja.FunctionCall) goja.Value {
//	return func(call goja.FunctionCall) goja.Value {
//		if format, ok := goja.AssertFunction(c.util.Get("format")); ok {
//			ret, err := format(c.util, call.Arguments...)
//			if err != nil {
//				panic(err)
//			}
//
//			p(ret.String())
//		} else {
//			panic(c.runtime.NewTypeError("util.format is not a function"))
//		}
//
//		return nil
//	}
//}

func Require(runtime *goja.Runtime, module *goja.Object) {
	requireWithPrinter()(runtime, module)
}

func RequireWithPrinter() require.ModuleLoader {
	return requireWithPrinter()
}

func requireWithPrinter() require.ModuleLoader {
	return func(runtime *goja.Runtime, module *goja.Object) {
		c := &TableEngine{
			runtime: runtime,
		}

		//c.util = require.Require(runtime, "util").(*goja.Object)

		o := module.Get("exports").(*goja.Object)
		o.Set("NewTableFieldClass", c.NewTableFieldClass)
		o.Set("NewTableFieldType", c.NewTableFieldType)
		o.Set("NewTableFieldEnumType", c.NewTableFieldEnumType)
		o.Set("NewTableFieldArrayType", c.NewTableFieldArrayType)
		o.Set("NewTableFieldMapType", c.NewTableFieldMapType)
		o.Set("NewTableFieldClassType", c.NewTableFieldClassType)
		o.Set("TableConfig", config.GlobalConfig.Table)
	}
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("tableEngine", require.Require(runtime, "tableEngine"))
}

func init() {
	require.RegisterNativeModule("tableEngine", Require)
}
