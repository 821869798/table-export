package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
	"github.com/821869798/table-export/meta"
)

type DMapType struct {
	valueType *meta.TableFieldType
	keyType   *meta.TableFieldType
	dataMap   map[apibase.DType]apibase.DType
}

func (dm *DMapType) String() string {
	result := ""
	for k, v := range dm.dataMap {
		result += fmt.Sprintf("%s:%s,", k.String(), v.String())
	}
	return fmt.Sprintf("<DMap>{%s}", result)
}

func (dm *DMapType) Value() map[apibase.DType]apibase.DType {
	return dm.dataMap
}

func (dm *DMapType) ValueType() *meta.TableFieldType {
	return dm.valueType
}

func (dm *DMapType) KeyType() *meta.TableFieldType {
	return dm.keyType
}

func (dm *DMapType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (dm *DMapType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (dm *DMapType) IsEmptyOrZero() bool {
	return len(dm.dataMap) == 0
}

func NewDMapType(keyType *meta.TableFieldType, valueType *meta.TableFieldType, data map[apibase.DType]apibase.DType) *DMapType {
	return &DMapType{
		keyType:   keyType,
		valueType: valueType,
		dataMap:   data,
	}
}
