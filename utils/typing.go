package utils

import (
	"strconv"
)

func AllAsInt[T int | int64 | uint](arr []string) []T {
	var newArr []T
	for _, ele := range arr {
		num, _ := strconv.Atoi(ele)
		newArr = append(newArr, T(num))
	}
	return newArr
}
