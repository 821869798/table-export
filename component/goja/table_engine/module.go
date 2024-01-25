package table_engine

import (
	"encoding/json"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/field_type"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	_ "github.com/dop251/goja_nodejs/util"
	"github.com/gookit/slog"
	"reflect"
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

func (c *TableEngine) forEach(jsValue goja.Value, callback goja.Callable) {
	// 从goja.Value中获取实际值
	rawValue := jsValue.Export()
	reflectValue := reflect.ValueOf(rawValue)

	// 检查是否是一个map
	switch reflectValue.Kind() {
	case reflect.Map:
		// 遍历map的所有键值对
		for _, key := range reflectValue.MapKeys() {
			val := reflectValue.MapIndex(key)

			// 将键和值转换为goja.Value
			jsKey := c.runtime.ToValue(key.Interface())
			jsVal := c.runtime.ToValue(val.Interface())

			// 调用JavaScript的回调函数
			_, err := callback(nil, jsKey, jsVal)
			if err != nil {
				slog.Fatal(err)
				return
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < reflectValue.Len(); i++ {
			elemValue := reflectValue.Index(i)
			jsElem := c.runtime.ToValue(elemValue.Interface())

			// 调用JavaScript的回调函数
			_, err := callback(nil, c.runtime.ToValue(i), jsElem)
			if err != nil {
				slog.Fatal(err)
				return
			}
		}
	default:
		slog.Fatalf("[forEach] unsupported type: %v", reflectValue.Kind())
	}
}

func (c *TableEngine) len(jsValue goja.Value) goja.Value {
	rawValue := jsValue.Export()
	reflectValue := reflect.ValueOf(rawValue)
	return c.runtime.ToValue(reflectValue.Len())
}

func (c *TableEngine) getMapValue(jsValue goja.Value, keyValue goja.Value) goja.Value {
	rawValue := jsValue.Export()
	mapValue := reflect.ValueOf(rawValue)
	if mapValue.Kind() != reflect.Map {
		slog.Error("param is not a map type")
		return goja.Undefined()
	}

	mapKeyType := mapValue.Type().Key()

	// 将 JavaScript 的键转换为 Go 中 map 的键类型
	var goKey reflect.Value
	switch mapKeyType.Kind() {
	case reflect.Int32, reflect.Int64:
		// 从 JavaScript 中获取的数字是 float64 类型，需要转换为 Go 的整型
		floatKey := keyValue.ToFloat()
		intKey := int64(floatKey)
		goKey = reflect.ValueOf(intKey).Convert(mapKeyType)
	case reflect.Uint32, reflect.Uint64:
		// 从 JavaScript 中获取的数字是 float64 类型，需要转换为 Go 的整型
		floatKey := keyValue.ToFloat()
		uintKey := uint64(floatKey)
		goKey = reflect.ValueOf(uintKey).Convert(mapKeyType)
	case reflect.Float32, reflect.Float64:
		goKey = reflect.ValueOf(keyValue.ToFloat()).Convert(mapKeyType)
	case reflect.Bool:
		goKey = reflect.ValueOf(keyValue.ToBoolean()).Convert(mapKeyType)
	case reflect.String:
		// 如果 Go 中的键是字符串，则直接使用它
		goKey = reflect.ValueOf(keyValue.String())
	default:
		fmt.Printf("Unsupported map key type: %v\n", mapKeyType.Kind())
		return goja.Undefined()
	}

	// 使用反射来查找对应的值
	reflectValue := mapValue.MapIndex(goKey)
	if !reflectValue.IsValid() {
		return goja.Undefined()
	}

	// 将找到的值转换为 goja.Value
	value := c.runtime.ToValue(reflectValue.Interface())
	return value
}

func (c *TableEngine) toJson(jsValue goja.Value) goja.Value {
	exportedValue := jsValue.Export()             // 将 goja.Value 转换为原生 Go 类型
	jsonBytes, err := json.Marshal(exportedValue) // 将 Go 对象序列化为 JSON
	if err != nil {
		slog.Errorf("Error marshaling to JSON:%v", err)
		return goja.Undefined()
	}
	return c.runtime.ToValue(string(jsonBytes))
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
		_ = o.Set("forEach", c.forEach)
		_ = o.Set("len", c.len)
		_ = o.Set("getMapValue", c.getMapValue)
		_ = o.Set("toJson", c.toJson)
	}
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("tableEngine", require.Require(runtime, "tableEngine"))
}

func init() {
	require.RegisterNativeModule("tableEngine", Require)
}
