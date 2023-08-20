package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
)

var (
	dlongPool []*DLongType
)

const dlongPoolSize = 128

func init() {
	dlongPool = make([]*DLongType, dlongPoolSize)
	for i := 0; i < dlongPoolSize; i++ {
		dlongPool[i] = &DLongType{
			value: int64(i),
		}
	}
}

type DLongType struct {
	value int64
}

func (dl *DLongType) String() string {
	return fmt.Sprintf("<DLong:%d>", dl.value)
}

func (dl *DLongType) Value() int64 {
	return dl.value
}

func (dl *DLongType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (dl *DLongType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (dl *DLongType) IsEmptyOrZero() bool {
	return dl.value == 0
}

func NewDLong(value int64) *DLongType {
	if value >= 0 && value < dlongPoolSize {
		return dlongPool[value]
	}
	return &DLongType{value: value}
}
