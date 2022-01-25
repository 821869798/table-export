package meta

import (
	"errors"
	"fmt"
	"regexp"
)

type TableField struct {
	Source     string
	Target     string
	Type       *TableFiledType
	TypeString string
	Desc       string
	Key        int
}

type TableFiledType struct {
	Type  FieldType       //类型
	Key   FieldType       //key的类型
	Value *TableFiledType //Value的类型
}

func newTableField(rtf *RawTableField) (*TableField, error) {
	tf := &TableField{
		Source:     rtf.Source,
		Target:     rtf.Target,
		TypeString: rtf.Type,
		Key:        rtf.Key,
	}
	tft, err := getFieldTypeFromString(tf.TypeString)
	if err != nil {
		return nil, err
	}
	tf.Type = tft
	return tf, nil
}

//是否可以作为map的key
func (tft *TableFiledType) IsBaseType() bool {
	_, ok := baseFiledType[tft.Type]
	return ok
}

func (tft *TableFiledType) IsComplexType() bool {
	switch tft.Type {
	case FieldType_Slice, FieldType_Map:
		return true
	}
	return false
}

func (tft *TableFiledType) GetKeyFieldType() (*TableFiledType, bool) {
	result, ok := baseFiledType[tft.Key]
	return result, ok
}

func getFieldTypeFromString(origin string) (*TableFiledType, error) {

	tft := &TableFiledType{
		Key:   FieldType_None,
		Value: nil,
	}
	switch origin {
	case "int":
		tft.Type = FieldType_Int
	case "bool":
		tft.Type = FieldType_Bool
	case "float":
		tft.Type = FiledType_Float
	case "string":
		tft.Type = FieldType_String
	default:
		reg := regexp.MustCompile(`^\[\](.+)$`)
		result := reg.FindAllStringSubmatch(origin, -1)
		if len(result) == 1 && len(result[0]) == 2 {
			tft.Type = FieldType_Slice
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
			tft.Type = FieldType_Map
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
