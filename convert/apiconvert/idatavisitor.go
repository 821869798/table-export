package apiconvert

import (
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/field_type"
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
	AcceptArray(r []interface{}, ValueType *field_type.TableFieldType)
	AcceptStringArray(r []string, ValueType *field_type.TableFieldType)
	AcceptStringMap(r map[string]string, KeyType *field_type.TableFieldType, ValueType *field_type.TableFieldType)
	AcceptMap(r map[string]interface{}, KeyType *field_type.TableFieldType, ValueType *field_type.TableFieldType)
	AcceptCommonMap(r map[interface{}]interface{}, KeyType *field_type.TableFieldType, ValueType *field_type.TableFieldType)
	AcceptClass(r map[string]interface{}, class *field_type.TableFieldClass)
	AcceptClassString(r map[string]string, class *field_type.TableFieldClass)
	AcceptClassNull(class *field_type.TableFieldClass)
}
