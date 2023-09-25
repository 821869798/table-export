package util

import "sort"

// GetMapSortedKeys 获取map的key的slice，并且排序
func GetMapSortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// 排序keys
	sort.Strings(keys)

	return keys
}
