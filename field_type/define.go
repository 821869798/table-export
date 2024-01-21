package field_type

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
	EFieldType_Int:    &TableFieldType{Type: EFieldType_Int},
	EFieldType_UInt:   &TableFieldType{Type: EFieldType_UInt},
	EFieldType_Long:   &TableFieldType{Type: EFieldType_Long},
	EFieldType_ULong:  &TableFieldType{Type: EFieldType_ULong},
	EFieldType_Bool:   &TableFieldType{Type: EFieldType_Bool},
	EFieldType_Float:  &TableFieldType{Type: EFieldType_Float},
	EFieldType_Double: &TableFieldType{Type: EFieldType_Double},
	EFieldType_String: &TableFieldType{Type: EFieldType_String},
	EFieldType_Enum:   &TableFieldType{Type: EFieldType_Enum},
}
