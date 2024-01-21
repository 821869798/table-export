package field_type

func NewTableFieldType(fieldType EFieldType) *TableFieldType {
	return newTableFieldType(fieldType)
}

func NewTableFieldEnumType(name string) *TableFieldType {
	return newTableFieldEnumType(name)
}

func NewTableFieldArrayType(value *TableFieldType) *TableFieldType {
	return newTableFieldArrayType(value)
}

func NewTableFieldMapType(key *TableFieldType, value *TableFieldType) *TableFieldType {
	return newTableFieldMapType(key, value)
}

func NewTableFieldClassType(class *TableFieldClass) *TableFieldType {
	return newTableFieldClassType(class)
}
