package ext_field

import (
	"errors"
	"fmt"
	"github.com/821869798/fankit/fanstr"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/field_type"
	"strconv"
)

type ExtFieldKVListIntFloat struct {
	ExtFieldKVListIntInt
}

func NewExtFieldKVListIntFloat() field_type.IExtFieldType {
	ft := field_type.NewTableFieldClass("KVList_IntFloat")
	ft.AddField("keys", field_type.NewTableFieldArrayType(field_type.NewTableFieldType(field_type.EFieldType_Int)))
	ft.AddField("values", field_type.NewTableFieldArrayType(field_type.NewTableFieldType(field_type.EFieldType_Float)))

	e := &ExtFieldKVListIntFloat{}
	e.FieldType = field_type.NewTableFieldClassType(ft)
	e.FieldType.SetExtFieldType(e)

	return e
}

func (e *ExtFieldKVListIntFloat) ParseOriginData(origin string) (interface{}, error) {
	strSlice := fanstr.SplitEx(origin, config.GlobalConfig.Table.MapSplit1)
	if len(strSlice) == 0 {
		return map[string]interface{}{
			"keys":   []interface{}{},
			"values": []interface{}{},
		}, nil
	}

	keys := make([]interface{}, 0, len(strSlice))
	values := make([]interface{}, 0, len(strSlice))
	for _, v := range strSlice {
		strSlice := fanstr.SplitEx(v, config.GlobalConfig.Table.MapSplit2)
		if len(strSlice) == 0 {
			continue
		}
		if len(strSlice) != 2 {
			return nil, errors.New(fmt.Sprintf("KVListIntFloat ParseOriginData error,need array length 2: [%v]", v))
		}
		k, err := strconv.Atoi(strSlice[0])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("KVListIntFloat ParseOriginData error,need param left int: [%v]", v))
		}
		v, err := strconv.ParseFloat(strSlice[1], 32)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("KVListIntFloat ParseOriginData error,need param right float: [%v]", v))
		}
		keys = append(keys, int32(k))
		values = append(values, float32(v))
	}

	return map[string]interface{}{
		"keys":   keys,
		"values": values,
	}, nil
}
