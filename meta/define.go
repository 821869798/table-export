package meta

type EFieldType int

const (
	FieldType_None EFieldType = iota
	FieldType_Int
	FieldType_UInt
	FieldType_Long
	FieldType_ULong
	FieldType_Bool
	FiledType_Float
	FiledType_Double
	FieldType_String
	FieldType_Slice
	FieldType_Map
)

var baseFiledType = map[EFieldType]*TableFieldType{
	FieldType_Int:    newTableFieldType(FieldType_Int),
	FieldType_UInt:   newTableFieldType(FieldType_UInt),
	FieldType_Long:   newTableFieldType(FieldType_Long),
	FieldType_ULong:  newTableFieldType(FieldType_ULong),
	FieldType_Bool:   newTableFieldType(FieldType_Bool),
	FiledType_Float:  newTableFieldType(FiledType_Float),
	FiledType_Double: newTableFieldType(FiledType_Double),
	FieldType_String: newTableFieldType(FieldType_String),
}
