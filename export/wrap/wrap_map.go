package wrap

import (
	"errors"
	"strings"
	"table-export/config"
	"table-export/export/common"
	"table-export/meta"
)

type mapWrap struct{}

func (b *mapWrap) OutputValue(exportType common.ExportType, filedType *meta.TableFiledType, origin string) (interface{}, error) {
	switch exportType {
	case common.ExportType_Lua:
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
	}
	return nil, nil
}
