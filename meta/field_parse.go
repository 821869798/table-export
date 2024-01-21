package meta

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/field_type"
	"regexp"
	"strings"
)

var (
	regArray       = regexp.MustCompile(`^\[\](.+)$`)
	regMap         = regexp.MustCompile(`^map\[(\w+)\](.+)$`)
	regWordGeneric = regexp.MustCompile(`^([\w_]+)<(.+)>$`)
)

// parseFieldTypeFromString 解析表格字段类型
func parseFieldTypeFromString(origin string) (*field_type.TableFieldType, error) {

	tft := field_type.NewTableFieldType(field_type.EFieldType_None)
	switch origin {
	case "int":
		tft.Type = field_type.EFieldType_Int
	case "uint":
		tft.Type = field_type.EFieldType_UInt
	case "long":
		tft.Type = field_type.EFieldType_Long
	case "ulong":
		tft.Type = field_type.EFieldType_ULong
	case "bool":
		tft.Type = field_type.EFieldType_Bool
	case "float":
		tft.Type = field_type.EFieldType_Float
	case "double":
		tft.Type = field_type.EFieldType_Double
	case "string":
		tft.Type = field_type.EFieldType_String
	default:

		// array
		result := regArray.FindAllStringSubmatch(origin, -1)
		if len(result) == 1 && len(result[0]) == 2 {
			tft.Type = field_type.EFieldType_Slice
			subTft, err := parseFieldTypeFromString(strings.TrimSpace(result[0][1]))
			if err != nil {
				return nil, err
			}
			tft.Value = subTft
			break
		}

		// map
		result = regMap.FindAllStringSubmatch(origin, -1)
		if len(result) == 1 && len(result[0]) == 3 {
			tft.Type = field_type.EFieldType_Map
			keyTft, err := parseFieldTypeFromString(strings.TrimSpace(result[0][1]))
			if err != nil {
				return nil, err
			}
			if !keyTft.IsBaseType() {
				return nil, errors.New(fmt.Sprintf("map key type not support[%v]", keyTft))
			}
			tft.Key = keyTft

			valueTft, err := parseFieldTypeFromString(result[0][2])
			if err != nil {
				return nil, err
			}
			tft.Value = valueTft
			break
		}

		// keyword generic: ext field_type type generic
		result = regWordGeneric.FindAllStringSubmatch(origin, -1)
		if len(result) == 1 && len(result[0]) == 3 {
			// TODO generic support
		}

		// ext field_type type
		extFieldType, ok := env.GetExtFieldType(origin)
		if ok {
			return extFieldType.TableFieldType(), nil
		}

		// enum
		enumDefine, ok := env.GetEnumDefine(origin)
		if ok {
			tft.Type = field_type.EFieldType_Enum
			tft.Name = enumDefine.Name
			break
		}

		tft = nil
	}

	if tft != nil {
		return tft, nil
	}
	return nil, errors.New("no support type:" + origin)
}
