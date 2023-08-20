package meta

type EFieldType int

const (
	EFieldType_None EFieldType = iota
	EFieldType_Int
	EFieldType_UInt
	EFieldType_Long
	EFieldType_ULong
	EFieldType_Bool
	EFiledType_Float
	EFiledType_Double
	EFieldType_String
	EFieldType_Slice
	EFieldType_Map
)

var baseFiledType = map[EFieldType]*TableFieldType{
	EFieldType_Int:    newTableFieldType(EFieldType_Int),
	EFieldType_UInt:   newTableFieldType(EFieldType_UInt),
	EFieldType_Long:   newTableFieldType(EFieldType_Long),
	EFieldType_ULong:  newTableFieldType(EFieldType_ULong),
	EFieldType_Bool:   newTableFieldType(EFieldType_Bool),
	EFiledType_Float:  newTableFieldType(EFiledType_Float),
	EFiledType_Double: newTableFieldType(EFiledType_Double),
	EFieldType_String: newTableFieldType(EFieldType_String),
}
