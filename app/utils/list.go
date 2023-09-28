package utils

import (
	"strconv"
	"strings"
)

func InSlice(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}

	return false
}

func ContainsSubSlice(str string, list []string) bool {
	for _, s := range list {
		if strings.Contains(s, str) {
			return true
		}
	}

	return false
}

func InIntSlice(str int, list []int) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}

	return false
}

func InInt64Slice(str int64, list []int64) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}

	return false
}

func RemoveRepByMap(slc []string) []string {
	result := []string{}         //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[strings.ToLower(e)] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func TransIntToStringSlice(intSlice []int64) []string {
	res := make([]string, 0)
	for _, intValue := range intSlice {
		res = append(res, strconv.FormatInt(intValue, 10))
	}
	return res
}

func RemoveFromSlice(str string, list []string) []string {
	list1 := list
	for k, v := range list {
		if v == str {
			list = append(list1[:k], list1[k+1:]...)
		}
	}
	return list
}

func SliceDiff(a, b []string) []string {
	var m = make(map[string]int)
	for _, item := range b {
		if _, ok := m[item]; !ok {
			m[item] = 1
		}
	}
	var res []string
	for _, item := range a {
		if _, ok := m[item]; !ok {
			res = append(res, item)
		}
	}
	return res
}

// ChunkStringArray 将数组切割指定长度的二维数组
func ChunkStringArray(arr []string, size int) [][]string {
	var result [][]string

	for i := 0; i < len(arr); i += size {
		end := i + size

		if end > len(arr) {
			end = len(arr)
		}

		result = append(result, arr[i:end])
	}

	return result
}
