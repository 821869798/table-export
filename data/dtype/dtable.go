package dtype

import (
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
)

type DTableType struct {
}

func (D DTableType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	//TODO implement me
	panic("implement me")
}

func (D DTableType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	//TODO implement me
	panic("implement me")
}

func (D DTableType) IsEmptyOrZero() bool {
	//TODO implement me
	panic("implement me")
}

func (D DTableType) String() string {
	//TODO implement me
	panic("implement me")
}
