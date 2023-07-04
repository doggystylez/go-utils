package utils

import "strings"

func Max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

func StringInSlice(str string, list []string) bool {
	for _, s := range list {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}

func ExclusiveAppend(slice []string, s string) []string {
	for _, element := range slice {
		if element == s {
			return slice
		}
	}
	return append(slice, s)
}

func ExclusiveCombine(slice []string, newSlice []string) []string {
	for _, newElement := range newSlice {
		exists := false
		for _, element := range slice {
			if element == newElement {
				exists = true
				break
			}
		}
		if !exists {
			slice = append(slice, newElement)
		}
	}
	return slice
}

func Exclude(slice []string, remove []string) (out []string) {
	toRemove := make(map[string]bool)
	for _, str := range remove {
		toRemove[str] = true
	}
	for _, str := range slice {
		if !toRemove[str] {
			out = append(out, str)
		}
	}
	return
}
