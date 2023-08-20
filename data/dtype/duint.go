package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
)

var (
	duintPool []*DUIntType
)

const duintPoolSize = 128

func init() {
	duintPool = make([]*DUIntType, duintPoolSize)
	for i := 0; i < duintPoolSize; i++ {
		duintPool[i] = &DUIntType{
			value: uint32(i),
		}
	}
}

type DUIntType struct {
	value uint32
}

func (du *DUIntType) String() string {
	return fmt.Sprintf("<DUInt:%d>", du.value)
}

func (du *DUIntType) Value() uint32 {
	return du.value
}

func (du *DUIntType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (du *DUIntType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (du *DUIntType) IsEmptyOrZero() bool {
	return du.value == 0
}

func NewDUInt(value uint32) *DUIntType {
	if value < duintPoolSize {
		return duintPool[value]
	}
	return &DUIntType{value: value}
}
