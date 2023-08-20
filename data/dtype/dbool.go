package dtype

import (
	"github.com/821869798/table-export/apibase"
	"github.com/821869798/table-export/data/check/target"
)

var (
	dboolDefault = &DBoolType{
		value: false,
	}
)

type DBoolType struct {
	value bool
}

func (db *DBoolType) String() string {
	if db.value {
		return "<DBool:false>"
	}
	return "<DBool:false>"
}

func (db *DBoolType) Value() bool {
	return db.value
}

func (db *DBoolType) AcceptVisitor(visitor apibase.IDTypeDataVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (db *DBoolType) AcceptCheckerVisitor(target *target.CheckerTarget, visitor apibase.ICheckerVisitor) {
	// TODO: Implement this method or return an error.
	panic("not implemented")
}

func (db *DBoolType) IsEmptyOrZero() bool {
	return !db.value
}

func NewDBool(value bool) *DBoolType {
	if !value {
		return dboolDefault
	}
	return &DBoolType{value: value}
}
