package adapter

import "github.com/821869798/table-export/meta"

type Map struct {
	Datas     map[string]string
	KeyType   *meta.TableFieldType
	ValueType *meta.TableFieldType
}

func NewMap(datas map[string]string, keyType *meta.TableFieldType, fieldType *meta.TableFieldType) *Map {
	m := &Map{
		Datas:     datas,
		KeyType:   keyType,
		ValueType: fieldType,
	}
	return m
}

var EmptyMap *Map = NewMap(nil, meta.TableFieldTypeNone, meta.TableFieldTypeNone)
