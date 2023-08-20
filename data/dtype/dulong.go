package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
)

var (
	dulongPool []*DULongType
)

const dulongPoolSize = 128

func init() {
	dulongPool = make([]*DULongType, dulongPoolSize)
	for i := 0; i < dulongPoolSize; i++ {
		dulongPool[i] = &DULongType{
			value: uint64(i),
		}
	}
}

type DULongType struct {
	value uint64
}

func (du *DULongType) String() string {
	return fmt.Sprintf("<DULong:%d>", du.value)
}

func (du *DULongType) Value() uint64 {
	return du.value
}

func (du *DULongType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (du *DULongType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (du *DULongType) IsEmptyOrZero() bool {
	return du.value == 0
}

func NewDULong(value uint64) *DULongType {
	if value < dulongPoolSize {
		return dulongPool[value]
	}
	return &DULongType{value: value}
}
