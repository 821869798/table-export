package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/field_type"
	"strings"
)

type mapWrap struct{}

func (b *mapWrap) OutputValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (interface{}, error) {
	origin = strings.TrimSpace(origin)
	switch exportType {
	case config.ExportType_Json:
		strMap := strings.Split(origin, config.GlobalConfig.Table.MapSplit1)
		result := make(map[string]interface{}, len(strMap))
		for _, v := range strMap {
			if v == "" {
				continue
			}
			kvStr := strings.Split(v, config.GlobalConfig.Table.MapSplit2)
			if len(kvStr) != 2 {
				return nil, errors.New("map format error to convert,string:" + origin)
			}

			keyContent, err := GetOutputStringValue(exportType, fieldType.Key, kvStr[0])
			if err != nil {
				return nil, err
			}

			//判断key是否重复
			_, ok := result[keyContent]
			if ok {
				return nil, errors.New("map format error to convert,key repeated,string:" + origin)
			}

			valueContent, err := GetOutputValue(exportType, fieldType.Value, kvStr[1])
			if err != nil {
				return nil, err
			}

			result[keyContent] = valueContent

		}
		return result, nil
	default:
		strMap := strings.Split(origin, config.GlobalConfig.Table.MapSplit1)
		result := make(map[interface{}]interface{}, len(strMap))
		for _, v := range strMap {
			if v == "" {
				continue
			}
			kvStr := strings.Split(v, config.GlobalConfig.Table.MapSplit2)
			if len(kvStr) != 2 {
				return nil, errors.New("map format error to convert,string:" + origin)
			}

			keyContent, err := GetOutputValue(exportType, fieldType.Key, kvStr[0])
			if err != nil {
				return nil, err
			}

			//判断key是否重复
			_, ok := result[keyContent]
			if ok {
				return nil, errors.New("map format error to convert,key repeated,string:" + origin)
			}

			valueContent, err := GetOutputValue(exportType, fieldType.Value, kvStr[1])
			if err != nil {
				return nil, err
			}

			result[keyContent] = valueContent

		}
		return result, nil
	}
}

func (b *mapWrap) OutputStringValue(exportType config.ExportType, fieldType *field_type.TableFieldType, origin string) (string, error) {
	origin = strings.TrimSpace(origin)
	switch exportType {
	case config.ExportType_Lua:
		strMap := strings.Split(origin, config.GlobalConfig.Table.MapSplit1)
		result := "{"
		//用来判断是否有key重复的情况
		keyMapping := make(map[string]bool)
		for _, v := range strMap {
			if v == "" {
				continue
			}
			kvStr := strings.Split(v, config.GlobalConfig.Table.MapSplit2)
			if len(kvStr) != 2 {
				return "", errors.New("map format error to convert,string:" + origin)
			}

			keyContent, err := GetOutputStringValue(exportType, fieldType.Key, kvStr[0])
			if err != nil {
				return "", err
			}
			//判断key是否重复
			_, ok := keyMapping[keyContent]
			if ok {
				return "", errors.New("map format error to convert,key repeated,string:" + origin)
			}
			keyMapping[keyContent] = true

			valueContent, err := GetOutputStringValue(exportType, fieldType.Value, kvStr[0])
			if err != nil {
				return "", err
			}

			result += "[" + keyContent + "]" + "=" + valueContent + ","

		}
		result += "}"
		return result, nil
	default:
		return "", errors.New(fmt.Sprintf("OutputStringValue map no support export type[%v]", exportType))
	}
}

func (b *mapWrap) OutputDefTypeValue(exportType config.ExportType, fieldType *field_type.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		keyDef, err := GetOutputDefTypeValue(exportType, fieldType.Key, collectionReadonly)
		if err != nil {
			return "", err
		}
		valueDef, err := GetOutputDefTypeValue(exportType, fieldType.Value, collectionReadonly)
		if err != nil {
			return "", err
		}
		var result string
		if collectionReadonly {
			result = fmt.Sprintf("IReadOnlyDictionary<%s, %s>", keyDef, valueDef)
		} else {
			result = fmt.Sprintf("Dictionary<%s, %s>", keyDef, valueDef)
		}

		return result, nil
	default:
		return "", errors.New("no support export Type Output DefType")
	}
}

func (b *mapWrap) FormatValueInterface(fieldType *field_type.TableFieldType, origin interface{}) (interface{}, error) {
	switch origin.(type) {
	case map[string]interface{}:
		return origin, nil
	case map[interface{}]interface{}:
		return origin, nil
	case map[string]string:
		return origin, nil
	default:
		return nil, errors.New(fmt.Sprintf("[FormatValueInterface|map] no support type[%T]", origin))
	}

}

func (b *mapWrap) DataVisitorString(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin string) error {
	origin = strings.TrimSpace(origin)
	if origin == "" {
		visitor.AcceptStringMap(EmptyStringMap, fieldType.Key, fieldType.Value)
		return nil
	}
	strMap := strings.Split(origin, config.GlobalConfig.Table.MapSplit1)
	result := make(map[string]string, len(strMap))
	for _, v := range strMap {
		if v == "" {
			continue
		}
		kvStr := strings.Split(v, config.GlobalConfig.Table.MapSplit2)
		if len(kvStr) != 2 {
			return errors.New("map format error to convert,string:" + origin)
		}

		_, ok := result[kvStr[0]]
		if ok {
			return errors.New("map format error to convert,key repeated,string:" + origin)
		}

		result[kvStr[0]] = kvStr[1]
	}
	visitor.AcceptStringMap(result, fieldType.Key, fieldType.Value)
	return nil
}

func (b *mapWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, fieldType *field_type.TableFieldType, origin interface{}) error {
	switch value := origin.(type) {
	case map[string]interface{}:
		visitor.AcceptMap(value, fieldType.Key, fieldType.Value)
		return nil
	case map[interface{}]interface{}:
		visitor.AcceptCommonMap(value, fieldType.Key, fieldType.Value)
		return nil
	case map[string]string:
		visitor.AcceptStringMap(value, fieldType.Key, fieldType.Value)
		return nil
	case string:
		return b.DataVisitorString(visitor, fieldType, value)
	default:
		return errors.New(fmt.Sprintf("[DataVisitorValue|map] no support type[%T]", origin))
	}
}

func (b *mapWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptMap(fieldType, fieldName, reader, depth)
}
