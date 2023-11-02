package apiconvert

import (
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/meta"
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
	AcceptArray(r []interface{}, ValueType *meta.TableFieldType)
	AcceptStringArray(r []string, ValueType *meta.TableFieldType)
	AcceptStringMap(r map[string]string, KeyType *meta.TableFieldType, ValueType *meta.TableFieldType)
	AcceptMap(r map[string]interface{}, KeyType *meta.TableFieldType, ValueType *meta.TableFieldType)
	AcceptCommonMap(r map[interface{}]interface{}, KeyType *meta.TableFieldType, ValueType *meta.TableFieldType)
}
