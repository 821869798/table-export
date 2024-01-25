package ext_field_script

import (
	"errors"
	"fmt"
	"github.com/821869798/fankit/fanpath"
	"github.com/821869798/table-export/component/goja/table_engine"
	"github.com/821869798/table-export/field_type"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"sync"
)

// ExtFieldJS 通过js脚本扩展字段类型
type ExtFieldJS struct {
	FieldType  *field_type.TableFieldType
	ScriptPath string
	jsResult   *jsScriptResult
	Error      string
	mutex      sync.Mutex
}

type jsScriptResult struct {
	Name           string
	DefineFile     string
	ParseFunc      func(string, *ExtFieldJS) interface{}
	TableFieldType func() *field_type.TableFieldType
}

func NewExtFieldJS(scripPath string) (field_type.IExtFieldType, error) {

	e := &ExtFieldJS{
		ScriptPath: scripPath,
		jsResult:   &jsScriptResult{},
	}
	registry := require.NewRegistry(require.WithGlobalFolders(fanpath.ExecuteParentPath()))

	vm := goja.New()
	_ = registry.Enable(vm)
	console.Enable(vm)
	table_engine.Enable(vm)

	var script = fmt.Sprintf("const extField = require(\"%s\");\n", e.ScriptPath) +
		`const result = { Name: extField.Name(),DefineFile: extField.DefineFile(),TableFieldType: extField.TableFieldType, ParseFunc: extField.ParseOriginData};
result;`
	res, err := vm.RunString(script)
	if err != nil {
		return nil, err
	}

	if err := vm.ExportTo(res, e.jsResult); err != nil {
		return nil, err
	}

	e.FieldType = e.jsResult.TableFieldType()
	e.FieldType.SetExtFieldType(e)

	return e, nil
}

func (e *ExtFieldJS) Name() string {
	return e.jsResult.Name
}

func (e *ExtFieldJS) DefineFile() string {
	return e.jsResult.DefineFile
}

func (e *ExtFieldJS) TableFieldType() *field_type.TableFieldType {
	return e.FieldType
}

func (e *ExtFieldJS) ParseOriginData(origin string) (interface{}, error) {
	// goja 不支持线程安全
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.Error = ""
	result := e.jsResult.ParseFunc(origin, e)
	if e.Error != "" {
		return nil, errors.New(e.Error)
	}
	return result, nil
}
