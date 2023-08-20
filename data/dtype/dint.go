package dtype

import (
	"fmt"
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
)

var (
	dintPool []*DIntType
)

const dintPoolSize = 128

func init() {
	dintPool = make([]*DIntType, dintPoolSize)
	for i := 0; i < dintPoolSize; i++ {
		dintPool[i] = &DIntType{
			value: int32(i),
		}
	}
}

type DIntType struct {
	value int32
}

func (di *DIntType) String() string {
	return fmt.Sprintf("<DInt:%d>", di.value)
}

func (di *DIntType) Value() int32 {
	return di.value
}

func (di *DIntType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (di *DIntType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (di *DIntType) IsEmptyOrZero() bool {
	return di.value == 0
}

func NewDInt(value int32) *DIntType {
	if value >= 0 && value < dintPoolSize {
		return dintPool[value]
	}
	return &DIntType{value: value}
}
