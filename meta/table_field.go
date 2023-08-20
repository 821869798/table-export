package meta

import (
	"errors"
	"fmt"
	"regexp"
)

type TableField struct {
	Source     string
	Target     string
	Type       *TableFieldType
	TypeString string
	Desc       string
	Key        int
}

func newTableField(rtf *RawTableField) (*TableField, error) {
	tf := &TableField{
		Source:     rtf.Source,
		Target:     rtf.Target,
		TypeString: rtf.Type,
		Desc:       rtf.Desc,
		Key:        rtf.Key,
	}
	tft, err := getFieldTypeFromString(tf.TypeString)
	if err != nil {
		return nil, err
	}
	tf.Type = tft
	return tf, nil
}

type TableFieldType struct {
	Type  EFieldType      //类型
	Key   EFieldType      //key的类型
	Value *TableFieldType //Value的类型
}

var TableFieldTypeNone = newTableFieldType(EFieldType_None)

func newTableFieldType(fieldType EFieldType) *TableFieldType {
	tft := &TableFieldType{
		Type:  fieldType,
		Key:   EFieldType_None,
		Value: nil,
	}
	return tft
}

func newTableFieldArrayType(value *TableFieldType) *TableFieldType {
	tft := &TableFieldType{
		Type:  EFieldType_Slice,
		Key:   EFieldType_None,
		Value: value,
	}
	return tft
}

func newTableFieldMapType(key EFieldType, value *TableFieldType) *TableFieldType {
	tft := &TableFieldType{
		Type:  EFieldType_Map,
		Key:   key,
		Value: value,
	}
	return tft
}

// 是否可以作为map的key
func (tft *TableFieldType) IsBaseType() bool {
	_, ok := baseFiledType[tft.Type]
	return ok
}

func (tft *TableFieldType) IsComplexType() bool {
	switch tft.Type {
	case EFieldType_Slice, EFieldType_Map:
		return true
	}
	return false
}

func (tft *TableFieldType) IsReferenceType() bool {
	if tft.IsComplexType() {
		return true
	}
	switch tft.Type {
	case EFieldType_String:
		return true
	}
	return false
}

func (tft *TableFieldType) GetKeyFieldType() (*TableFieldType, bool) {
	result, ok := baseFiledType[tft.Key]
	return result, ok
}

func (tft *TableFieldType) CreateArrayFieldType() *TableFieldType {
	var fieldType = newTableFieldArrayType(tft)
	return fieldType
}

func getFieldTypeFromString(origin string) (*TableFieldType, error) {

	tft := newTableFieldType(EFieldType_None)
	switch origin {
	case "int":
		tft.Type = EFieldType_Int
	case "uint":
		tft.Type = EFieldType_UInt
	case "long":
		tft.Type = EFieldType_Long
	case "ulong":
		tft.Type = EFieldType_ULong
	case "bool":
		tft.Type = EFieldType_Bool
	case "float":
		tft.Type = EFiledType_Float
	case "double":
		tft.Type = EFiledType_Double
	case "string":
		tft.Type = EFieldType_String
	default:
		reg := regexp.MustCompile(`^\[\](.+)$`)
		result := reg.FindAllStringSubmatch(origin, -1)
		if len(result) == 1 && len(result[0]) == 2 {
			tft.Type = EFieldType_Slice
			subTft, err := getFieldTypeFromString(result[0][1])
			if err != nil {
				return nil, err
			}
			tft.Value = subTft
			break
		}

		// map
		reg = regexp.MustCompile(`^map\[(\w+)\](.+)$`)
		result = reg.FindAllStringSubmatch(origin, -1)
		if len(result) == 1 && len(result[0]) == 3 {
			tft.Type = EFieldType_Map
			keyTft, err := getFieldTypeFromString(result[0][1])
			if err != nil {
				return nil, err
			}
			if !keyTft.IsBaseType() {
				return nil, errors.New(fmt.Sprintf("map key type not support[%v]", keyTft))
			}
			tft.Key = keyTft.Type

			valueTft, err := getFieldTypeFromString(result[0][2])
			if err != nil {
				return nil, err
			}
			tft.Value = valueTft
			break
		}

		tft = nil
	}

	if tft != nil {
		return tft, nil
	}
	return nil, errors.New("no support type:" + origin)
}
