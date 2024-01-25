package table_engine

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/field_type"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	_ "github.com/dop251/goja_nodejs/util"
	"github.com/gookit/slog"
)

type TableEngine struct {
	runtime *goja.Runtime
	//util        *goja.Object
}

func (c *TableEngine) Fatal(msg string) {
	slog.Fatal(msg)
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

func (c *TableEngine) GetExtFieldType(name string) *field_type.TableFieldType {
	v, ok := env.GetExtFieldType(name)
	if ok {
		return v.TableFieldType()
	}
	return nil
}

func (c *TableEngine) ParseExtFieldValue(name string, origin string) interface{} {
	v, ok := env.GetExtFieldType(name)
	if !ok {
		slog.Fatalf("parse ext field type error: name[%v],origin[%v]", name, origin)
	}
	result, err := v.ParseOriginData(origin)
	if err != nil {
		slog.Fatalf("parse ext field type error: name[%v],origin[%v],err[%v]", name, origin, err)
	}
	return result
}

func Require(runtime *goja.Runtime, module *goja.Object) {
	requireWithPrinter()(runtime, module)
}

func requireWithPrinter() require.ModuleLoader {
	return func(runtime *goja.Runtime, module *goja.Object) {
		c := &TableEngine{
			runtime: runtime,
		}

		//c.util = require.Require(runtime, "util").(*goja.Object)

		o := module.Get("exports").(*goja.Object)
		_ = o.Set("NewTableFieldClass", c.NewTableFieldClass)
		_ = o.Set("NewTableFieldType", c.NewTableFieldType)
		_ = o.Set("NewTableFieldEnumType", c.NewTableFieldEnumType)
		_ = o.Set("NewTableFieldArrayType", c.NewTableFieldArrayType)
		_ = o.Set("NewTableFieldMapType", c.NewTableFieldMapType)
		_ = o.Set("NewTableFieldClassType", c.NewTableFieldClassType)
		_ = o.Set("TableConfig", config.GlobalConfig.Table)
		_ = o.Set("Fatal", c.Fatal)
		_ = o.Set("GetExtFieldType", c.GetExtFieldType)
		_ = o.Set("ParseExtFieldValue", c.ParseExtFieldValue)

	}
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("tableEngine", require.Require(runtime, "tableEngine"))
}

func init() {
	require.RegisterNativeModule("tableEngine", Require)
}
