package ext_field

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/field_type"
	"github.com/821869798/table-export/util"
	"strconv"
)

// ExtFieldPointInt 示例，使用go扩展一个字段类型
type ExtFieldPointInt struct {
	FieldType *field_type.TableFieldType
}

func NewExtFieldPointInt() field_type.IExtFieldType {
	ft := field_type.NewTableFieldClass("PointInt")
	ft.AddField("x", field_type.NewTableFieldType(field_type.EFieldType_Int))
	ft.AddField("y", field_type.NewTableFieldType(field_type.EFieldType_Int))

	e := &ExtFieldPointInt{
		FieldType: field_type.NewTableFieldClassType(ft),
	}

	e.FieldType.SetExtFieldType(e)

	return e
}

func (e *ExtFieldPointInt) Name() string {
	return e.FieldType.Name
}

func (e *ExtFieldPointInt) DefineFile() string {
	return "CfgMath"
}

func (e *ExtFieldPointInt) TableFieldType() *field_type.TableFieldType {
	return e.FieldType
}

func (e *ExtFieldPointInt) ParseDataOne(origin string) (interface{}, error) {
	strSlice := util.SplitEx(origin, config.GlobalConfig.Table.ArraySplit)
	if len(strSlice) == 0 {
		return map[string]interface{}{
			"x": int32(0),
			"y": int32(0),
		}, nil
	}
	if len(strSlice) != 2 {
		return nil, errors.New(fmt.Sprintf("PointInt ParseDataOne error,need array length 2 or 0: [%v]", origin))
	}

	x, err := strconv.Atoi(strSlice[0])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("PointInt ParseDataOne error,need param x int: [%v]", origin))
	}
	y, err := strconv.Atoi(strSlice[1])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("PointInt ParseDataOne error,need param y int: [%v]", origin))
	}
	return map[string]interface{}{
		"x": int32(x),
		"y": int32(y),
	}, nil
}

func (e *ExtFieldPointInt) ParseDataMultiple(originArray []string) (interface{}, error) {
	return nil, errors.New("PointInt no support ParseDataMultiple")
}
