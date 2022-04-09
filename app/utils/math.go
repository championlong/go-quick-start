package utils

import (
	"fmt"
	"strconv"
)

func DivideFloat(a float64, b float64) float64 {
	if b > 0 {
		return Fixed4(a / b)
	}
	return 0
}

func DivideFloatWithInt(a float64, b int) float64 {
	if b > 0 {
		return Fixed4(a / float64(b))
	}
	return 0
}

func DivideInt(a int, b int) float64 {
	if b > 0 {
		return Fixed4(float64(a) / float64(b))
	}
	return 0
}

func DivideIntWithFloat(a int, b float64) float64 {
	if b > 0 {
		return Fixed4(float64(a) / b)
	}
	return 0
}

func Fixed4(f float64) float64 {
	return Fixed(f, 4)
}

func Fixed2(f float64) float64 {
	return Fixed(f, 2)
}

//保留precision位小数
func Fixed(f float64, precision int) float64 {
	if f == 0 {
		return 0
	}
	floatStr := fmt.Sprintf("%."+strconv.Itoa(precision)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

//计算增长率
func CalcIntGrowthPercentage(a1, a2 int) float64 {
	return CalcGrowthPercentage(float64(a1), float64(a2))
}

//计算增长率
func CalcGrowthPercentage(a1, a2 float64) float64 {
	if a2 > 0 {
		return Fixed4((a1 - a2) * 100 / a2)
	} else {
		return 0
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//返回第一个不为0的int
func CoalesceFloat64(list ...float64) float64 {
	for _, i := range list {
		if i != 0 {
			return i
		}
	}
	return 0
}

func StringSliceIn(element string, sliceInfo []string) bool {
	for _, item := range sliceInfo {
		if  item == element {
			return true
		}
	}
	return false
}