package util

import (
	"fmt"
	"strings"
)

// FormatPascalString 转换成帕斯卡命名法，每个单词首字母大写
func FormatPascalString(name string) string {
	result := ""
	strs := strings.Split(name, "_")
	for _, str := range strs {
		if str == "" {
			continue
		}

		result += Capitalize(str)
	}
	return result
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// ReplaceLast 从后往替换
func ReplaceLast(source string, strToReplace string, strWithReplace string) string {
	index := strings.LastIndex(source, strToReplace)
	if index == -1 {
		return source
	}
	return source[:index] + strWithReplace + source[index+len(strToReplace):]
}

// ReplaceWindowsLineEnd 更换行尾符\r\n为\n
func ReplaceWindowsLineEnd(source string) string {
	return strings.ReplaceAll(source, "\r\n", "\n")
}

func JoinEx[T fmt.Stringer](joinStrings []T, sep string) string {
	result := make([]string, 0, len(joinStrings))
	for _, str := range joinStrings {
		result = append(result, str.String())
	}

	return strings.Join(result, sep)
}
