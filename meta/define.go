package meta

type EFieldType int

const (
	EFieldType_None EFieldType = iota
	EFieldType_Int
	EFieldType_UInt
	EFieldType_Long
	EFieldType_ULong
	EFieldType_Bool
	EFieldType_Float
	EFieldType_Double
	EFieldType_String
	EFieldType_Enum
	EFieldType_Slice
	EFieldType_Map
	EFieldType_Class
)

var baseFieldType = map[EFieldType]*TableFieldType{
	EFieldType_Int:    newTableFieldType(EFieldType_Int),
	EFieldType_UInt:   newTableFieldType(EFieldType_UInt),
	EFieldType_Long:   newTableFieldType(EFieldType_Long),
	EFieldType_ULong:  newTableFieldType(EFieldType_ULong),
	EFieldType_Bool:   newTableFieldType(EFieldType_Bool),
	EFieldType_Float:  newTableFieldType(EFieldType_Float),
	EFieldType_Double: newTableFieldType(EFieldType_Double),
	EFieldType_String: newTableFieldType(EFieldType_String),
	EFieldType_Enum:   newTableFieldType(EFieldType_Enum),
}
