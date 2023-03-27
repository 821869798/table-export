package api

import (
	"table-export/convert/adapter"
	"table-export/data/model"
)

type IDataVisitor interface {
	AcceptTable(dataModel *model.TableModel)
	AcceptInt(r int32)
	AcceptUInt(r uint32)
	AcceptLong(r int64)
	AcceptULong(r uint64)
	AcceptBool(r bool)
	AcceptFloat(r float32)
	AcceptDouble(r float64)
	AcceptString(r string)
	AcceptByte(r byte)
	AcceptArray(r *adapter.Array)
	AcceptMap(r *adapter.Map)
}
