package meta

import (
	"errors"
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"strconv"
)

type TableCheck struct {
	Code    string
	CodeVM  *vm.Program
	CodeEnv map[string]interface{}
}

func newTableCheck(check *RawTableCheck) *TableCheck {
	return &TableCheck{
		Code: check.Code,
	}
}

func (c *TableCheck) CompileCode(memTable map[interface{}]interface{}, global map[string]map[interface{}]interface{}) error {

	c.CodeEnv = map[string]interface{}{
		//"$":      0,
		"table":  memTable,
		"global": global,
	}

	args := []expr.Option{
		expr.Env(c.CodeEnv),
		expr.AsBool(),
		expr.AllowUndefinedVariables(),
	}

	args = append(args, CustomBuiltinFunctions...)

	codeVM, err := expr.Compile(c.Code, args...)
	c.CodeVM = codeVM
	return err
}

func (c *TableCheck) Run(record map[string]interface{}) (bool, error) {
	c.CodeEnv["$"] = record

	resultAny, err := expr.Run(c.CodeVM, c.CodeEnv)
	if err != nil {
		return false, err
	}

	result, ok := resultAny.(bool)
	if !ok {
		return false, errors.New(fmt.Sprintf("table check result is not bool, code[%s] result[%v]", c.Code, resultAny))
	}
	return result, nil
}

func (c *TableCheck) RunGlobal() (bool, error) {
	resultAny, err := expr.Run(c.CodeVM, c.CodeEnv)
	if err != nil {
		return false, err
	}

	result, ok := resultAny.(bool)
	if !ok {
		return false, errors.New(fmt.Sprintf("table check result is not bool, code[%s] result[%v]", c.Code, resultAny))
	}
	return result, nil
}

var (
	CustomBuiltinFunctions = []expr.Option{
		expr.Function("int32", func(args ...interface{}) (interface{}, error) {
			switch value := args[0].(type) {
			case int:
				return int32(value), nil
			case string:
				num, err := strconv.Atoi(value)
				return int32(num), err
			case int32:
				return value, nil
			case float32:
				return int32(value), nil
			case uint32:
				return int32(value), nil
			case int64:
				return int32(value), nil
			case float64:
				return int32(value), nil
			case uint64:
				return int32(value), nil
			default:
				return nil, errors.New(fmt.Sprintf("int32 no support type[%T]", args[0]))
			}
		}),
		expr.Function("int64", func(args ...interface{}) (interface{}, error) {
			switch value := args[0].(type) {
			case int:
				return int64(value), nil
			case string:
				num, err := strconv.Atoi(value)
				return int64(num), err
			case int32:
				return int64(value), nil
			case float32:
				return int64(value), nil
			case uint32:
				return int64(value), nil
			case int64:
				return value, nil
			case float64:
				return int64(value), nil
			case uint64:
				return int64(value), nil
			default:
				return nil, errors.New(fmt.Sprintf("int64 no support type[%T]", args[0]))
			}
		}),
		expr.Function("uint32", func(args ...interface{}) (interface{}, error) {
			switch value := args[0].(type) {
			case int:
				return uint32(value), nil
			case string:
				num, err := strconv.Atoi(value)
				return uint32(num), err
			case int32:
				return uint32(value), nil
			case float32:
				return uint32(value), nil
			case uint32:
				return value, nil
			case int64:
				return uint32(value), nil
			case float64:
				return uint32(value), nil
			case uint64:
				return uint32(value), nil
			default:
				return nil, errors.New(fmt.Sprintf("uint32 no support type[%T]", args[0]))
			}
		}),
		expr.Function("uint64", func(args ...interface{}) (interface{}, error) {
			switch value := args[0].(type) {
			case int:
				return uint64(value), nil
			case string:
				num, err := strconv.Atoi(value)
				return uint64(num), err
			case int32:
				return uint64(value), nil
			case float32:
				return uint64(value), nil
			case uint32:
				return uint64(value), nil
			case int64:
				return uint64(value), nil
			case float64:
				return uint64(value), nil
			case uint64:
				return value, nil
			default:
				return nil, errors.New(fmt.Sprintf("uint64 no support type[%T]", args[0]))
			}
		}),
		expr.Function("float32", func(args ...interface{}) (interface{}, error) {
			switch value := args[0].(type) {
			case float64:
				return float32(value), nil
			case int:
				return float32(value), nil
			case string:
				num, err := strconv.ParseFloat(value, 32)
				return float32(num), err
			case int32:
				return float32(value), nil
			case float32:
				return value, nil
			case uint32:
				return float32(value), nil
			case int64:
				return float32(value), nil
			default:
				return nil, errors.New(fmt.Sprintf("uint64 no support type[%T]", args[0]))
			}
		}),
		expr.Function("float64", func(args ...interface{}) (interface{}, error) {
			switch value := args[0].(type) {
			case float64:
				return value, nil
			case int:
				return float64(value), nil
			case string:
				num, err := strconv.ParseFloat(value, 64)
				return num, err
			case int32:
				return float64(value), nil
			case float32:
				return float64(value), nil
			case uint32:
				return float64(value), nil
			case int64:
				return float64(value), nil
			default:
				return nil, errors.New(fmt.Sprintf("uint64 no support type[%T]", args[0]))
			}
		}),
	}
)
