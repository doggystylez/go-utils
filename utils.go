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

func WeightedAverage(values, weights []float64) float64 {
	var sum, totalWeight float64
	if len(values) != len(weights) {
		panic("number of values must equal number of weights")
	}
	if len(values) == 0 {
		return 0
	}
	for i, val := range values {
		sum += val * weights[i]
		totalWeight += weights[i]
	}
	return sum / totalWeight
}

func SubStringInSlice(str string, list []string) bool {
	for _, s := range list {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}

func StringInSlice(str string, list []string) bool {
	for _, s := range list {
		if s == str {
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
