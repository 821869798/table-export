package wrap

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/adapter"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/meta"
	"strings"
)

type mapWrap struct{}

func (b *mapWrap) OutputValue(exportType config.ExportType, filedType *meta.TableFieldType, origin string) (interface{}, error) {
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
				return nil, errors.New("map format error to convert,string:" + origin)
			}
			keyType, ok := filedType.GetKeyFieldType()
			if !ok {
				return nil, errors.New("map get Key field type error")
			}

			keyContent, err := GetOutputValue(exportType, keyType, kvStr[0])
			if err != nil {
				return nil, err
			}
			keyStr, ok := keyContent.(string)
			if !ok {
				return nil, errors.New("map get output value error")
			}

			//判断key是否重复
			_, ok = keyMapping[keyStr]
			if ok {
				return nil, errors.New("map format error to convert,key repeated,string:" + origin)
			}
			keyMapping[keyStr] = true

			valueContent, err := GetOutputValue(exportType, filedType.Value, kvStr[0])
			if err != nil {
				return nil, err
			}
			valueStr, ok := valueContent.(string)
			if !ok {
				return nil, errors.New("map get output value error")
			}

			result += "[" + keyStr + "]" + "=" + valueStr + ","

		}
		result += "}"
		return result, nil
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
			keyType, ok := filedType.GetKeyFieldType()
			if !ok {
				return nil, errors.New("map get Key field type error")
			}

			keyContent, err := GetOutputValue(exportType, keyType, kvStr[0])
			if err != nil {
				return nil, err
			}

			formatKey, err := GetOutputStringValue(exportType, keyType, keyContent)
			if err != nil {
				return nil, err
			}

			//判断key是否重复
			_, ok = result[formatKey]
			if ok {
				return nil, errors.New("map format error to convert,key repeated,string:" + origin)
			}

			valueContent, err := GetOutputValue(exportType, filedType.Value, kvStr[1])
			if err != nil {
				return nil, err
			}

			result[formatKey] = valueContent

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
			keyType, ok := filedType.GetKeyFieldType()
			if !ok {
				return nil, errors.New("map get Key field type error")
			}

			keyContent, err := GetOutputValue(exportType, keyType, kvStr[0])
			if err != nil {
				return nil, err
			}

			//判断key是否重复
			_, ok = result[keyContent]
			if ok {
				return nil, errors.New("map format error to convert,key repeated,string:" + origin)
			}

			valueContent, err := GetOutputValue(exportType, filedType.Value, kvStr[1])
			if err != nil {
				return nil, err
			}

			result[keyContent] = valueContent

		}
		return result, nil
	}
}

func (b *mapWrap) OutputStringValue(exportType config.ExportType, filedType *meta.TableFieldType, origin interface{}) (string, error) {
	return "", errors.New("map no support OutputStringValue")
}

func (b *mapWrap) OutputDefTypeValue(exportType config.ExportType, filedType *meta.TableFieldType, collectionReadonly bool) (string, error) {
	switch exportType {
	case config.ExportType_CS_Bin:
		keyType, _ := filedType.GetKeyFieldType()
		keyDef, err := GetOutputDefTypeValue(exportType, keyType, collectionReadonly)
		if err != nil {
			return "", err
		}
		valueDef, err := GetOutputDefTypeValue(exportType, filedType.Value, collectionReadonly)
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
	}
	return "", errors.New("no support export Type Output DefType")
}

func (b *mapWrap) DataVisitorValue(visitor apiconvert.IDataVisitor, filedType *meta.TableFieldType, origin string) error {
	if origin == "" {
		visitor.AcceptMap(adapter.EmptyMap)
		return nil
	}
	strMap := strings.Split(origin, config.GlobalConfig.Table.MapSplit1)
	result := make(map[string]string, len(strMap))
	keyType, ok := filedType.GetKeyFieldType()
	if !ok {
		return errors.New("map get Key field type error")
	}
	for _, v := range strMap {
		if v == "" {
			continue
		}
		kvStr := strings.Split(v, config.GlobalConfig.Table.MapSplit2)
		if len(kvStr) != 2 {
			return errors.New("map format error to convert,string:" + origin)
		}

		_, ok = result[kvStr[0]]
		if ok {
			return errors.New("map format error to convert,key repeated,string:" + origin)
		}

		result[kvStr[0]] = kvStr[1]
	}
	m := adapter.NewMap(result, keyType, filedType.Value)
	visitor.AcceptMap(m)
	return nil
}

func (b *mapWrap) CodePrintValue(print apiconvert.ICodePrinter, fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return print.AcceptMap(fieldType, fieldName, reader, depth)
}
