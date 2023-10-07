package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func StringListToStringBoolMap(arr []string) (result map[string]bool) {
	result = make(map[string]bool)
	for _, v := range arr {
		if v = strings.TrimSpace(v); v != "" {
			result[v] = true
		}
	}
	return
}

func IsInStrArray(src string, arr []string) bool {
	for _, element := range arr {
		if src == element {
			return true
		}
	}
	return false
}

func FirstStrIsCapital(str string) bool {
	if str == "" {
		return false
	}
	arr := strings.Split(str, "")
	tmp := arr[0]
	if arr[0] == strings.ToUpper(tmp) {
		return true
	} else {
		return false
	}
}

func IntList2StringList(is []int) []string {
	ret := make([]string, len(is))
	for i, v := range is {
		ret[i] = strconv.Itoa(v)
	}
	return ret
}

func StringContainsSubstrArray(s string, substr []string) bool {
	for _, sub := range substr {
		return strings.Contains(s, sub)
	}
	return false
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func IsNumber(str string) bool {
	for _, val := range str {
		if !unicode.IsNumber(val) {
			return false
		}
	}
	return true
}
