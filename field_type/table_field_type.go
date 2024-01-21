package field_type

type TableFieldType struct {
	Type         EFieldType       //类型
	Key          *TableFieldType  //key的类型
	Value        *TableFieldType  //Value的类型
	Name         string           // 枚举或者Class的名字
	Class        *TableFieldClass // 自定义类
	ExtFieldType IExtFieldType    // 扩展自定义字段类型
}

var TableFieldTypeNone = newTableFieldType(EFieldType_None)

func newTableFieldType(fieldType EFieldType) *TableFieldType {
	if fieldType != EFieldType_Enum {
		tft, ok := baseFieldType[fieldType]
		if ok {
			return tft
		}
	}
	tft := &TableFieldType{
		Type:  fieldType,
		Key:   nil,
		Value: nil,
	}
	return tft
}

func newTableFieldEnumType(name string) *TableFieldType {
	tft := &TableFieldType{
		Type: EFieldType_Enum,
		Name: name,
	}
	return tft
}

func newTableFieldArrayType(value *TableFieldType) *TableFieldType {
	tft := &TableFieldType{
		Type:  EFieldType_Slice,
		Key:   nil,
		Value: value,
	}
	return tft
}

func newTableFieldMapType(key *TableFieldType, value *TableFieldType) *TableFieldType {
	tft := &TableFieldType{
		Type:  EFieldType_Map,
		Key:   key,
		Value: value,
	}
	return tft
}

func newTableFieldClassType(class *TableFieldClass) *TableFieldType {
	tft := &TableFieldType{
		Type:  EFieldType_Class,
		Class: class,
		Name:  class.Name,
	}
	return tft
}

// SetExtFieldType 设置扩展字段类型
func (tft *TableFieldType) SetExtFieldType(extFieldType IExtFieldType) {
	tft.ExtFieldType = extFieldType
}

// IsBaseType 是否可以作为map的key
func (tft *TableFieldType) IsBaseType() bool {
	_, ok := baseFieldType[tft.Type]
	return ok
}

func (tft *TableFieldType) IsComplexType() bool {
	switch tft.Type {
	case EFieldType_Slice, EFieldType_Map, EFieldType_Class:
		return true
	default:
		return false
	}
}

func (tft *TableFieldType) IsReferenceType() bool {
	if tft.IsComplexType() {
		return true
	}
	if tft.Type == EFieldType_String {
		return true
	}
	return false
}

func (tft *TableFieldType) CreateArrayFieldType() *TableFieldType {
	var fieldType = newTableFieldArrayType(tft)
	return fieldType
}
