package array

import (
	"fmt"
	"strings"
)

func InArray(val int, array []int) bool {
	for _, i := range array {
		if val == i {
			return true
		}
	}
	return false
}

func InArray64(val int64, array []int64) bool {
	for _, i := range array {
		if val == i {
			return true
		}
	}
	return false
}

func InArrayString(val string, array []string) bool {
	for _, i := range array {
		if val == i {
			return true
		}
	}
	return false
}

func ArrayToString(array []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(array), " ", delim, -1), "[]")
}
