package meta

type FieldType int

const (
	FieldType_None FieldType = iota
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

var baseFiledType = map[FieldType]*TableFiledType{
	FieldType_Int:    &TableFiledType{Type: FieldType_Int},
	FieldType_UInt:   &TableFiledType{Type: FieldType_UInt},
	FieldType_Long:   &TableFiledType{Type: FieldType_Long},
	FieldType_ULong:  &TableFiledType{Type: FieldType_ULong},
	FieldType_Bool:   &TableFiledType{Type: FieldType_Bool},
	FiledType_Float:  &TableFiledType{Type: FiledType_Float},
	FiledType_Double: &TableFiledType{Type: FiledType_Double},
	FieldType_String: &TableFiledType{Type: FieldType_String},
}
