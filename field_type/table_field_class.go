package field_type

type TableFieldClass struct {
	Name         string
	fields       []TableFieldClassField
	fieldMapping map[string]TableFieldClassField
}

type TableFieldClassField struct {
	Name string
	Type *TableFieldType
}

func NewTableFieldClass(name string) *TableFieldClass {
	return &TableFieldClass{
		Name:         name,
		fieldMapping: make(map[string]TableFieldClassField),
	}
}

func (c *TableFieldClass) AddField(name string, fieldType *TableFieldType) {
	field := TableFieldClassField{
		Name: name,
		Type: fieldType,
	}
	c.fields = append(c.fields, field)
	c.fieldMapping[name] = field
}

func (c *TableFieldClass) AllFields() []TableFieldClassField {
	return c.fields
}
