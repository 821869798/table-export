package visitor

import (
	"cmp"
	"github.com/821869798/table-export/field_type"
	"github.com/gookit/slog"
	"slices"
)

// GetMapSortedKeys 获取map的key的slice，并且排序
func GetMapSortedKeys[T cmp.Ordered, V any](m map[T]V) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// 排序keys
	slices.Sort(keys)
	return keys
}

func GetMapSortedKeysInterface(m map[interface{}]interface{}, KeyType *field_type.TableFieldType) []interface{} {
	keys := make([]interface{}, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// 排序keys
	slices.SortFunc(keys, func(a, b interface{}) int {
		switch value := a.(type) {
		case int32:
			return cmp.Compare(value, b.(int32))
		case float32:
			return cmp.Compare(value, b.(float32))
		case string:
			return cmp.Compare(value, b.(string))
		case uint32:
			return cmp.Compare(value, b.(uint32))
		case int64:
			return cmp.Compare(value, b.(int64))
		case float64:
			return cmp.Compare(value, b.(float64))
		case uint64:
			return cmp.Compare(value, b.(uint64))
		case int:
			return cmp.Compare(value, b.(int))
		case bool:
			if value == b.(bool) {
				return 0
			} else if value {
				return 1
			} else {
				return -1
			}
		default:
			slog.Fatalf("unsupported type: %T", value)
			return 0
		}
	})
	return keys
}
