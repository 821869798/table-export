package apibase

import (
	"fmt"
	"github.com/821869798/table-export/data/check/target"
)

type DType interface {
	AcceptVisitor(visitor IDTypeDataVisitor)
	AcceptCheckerVisitor(target *target.CheckerTarget, visitor ICheckerVisitor)
	IsEmptyOrZero() bool
	fmt.Stringer
}
