package ext_post

import (
	"errors"
	"fmt"
	"github.com/821869798/fankit/fanpath"
	"github.com/821869798/table-export/component/goja/table_engine"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/ext/apiext"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

type ExtPostGlobalJS struct {
	ScriptPath string
	Error      string
	jsFunc     func(tableMap map[string]*model.TableModel, context *ExtPostGlobalJS)
}

func NewExtPostGlobalJS(scriptPath string) (apiext.IExtPostGlobal, error) {
	e := &ExtPostGlobalJS{
		ScriptPath: scriptPath,
	}

	registry := require.NewRegistry(require.WithGlobalFolders(fanpath.ExecuteParentPath()))

	vm := goja.New()
	_ = registry.Enable(vm)
	console.Enable(vm)
	table_engine.Enable(vm)

	var script = fmt.Sprintf("const postGlobal = require(\"%s\");\npostGlobal;", e.ScriptPath)

	res, err := vm.RunString(script)
	if err != nil {
		return nil, err
	}

	if err := vm.ExportTo(res, &e.jsFunc); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *ExtPostGlobalJS) PostGlobal(tableMap map[string]*model.TableModel) error {
	e.Error = ""
	e.jsFunc(tableMap, e)
	if e.Error != "" {
		return errors.New(e.Error)
	}
	return nil
}
