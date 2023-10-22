package math

// Max returns the maximum value in a slice of ints.
func Max(values []int) int {
	if len(values) == 0 {
		return 0
	}
	result := values[0]
	for _, num := range values {
		if num > result {
			result = num
		}
	}
	return result
}

// Min returns the minimum value in a slice of ints.
func Min(values []int) int {
	if len(values) == 0 {
		return 0
	}
	result := values[0]
	for _, num := range values {
		if num < result {
			result = num
		}
	}
	return result
}

// Sum returns the sum of a slice of ints.
func Sum(values []int) int {
	var result int
	for _, num := range values {
		result += num
	}
	return result
}

// Average returns the average of a slice of ints.
func Average(values []int) int {
	if len(values) == 0 {
		return 0
	}
	var sum int
	for _, num := range values {
		sum += num
	}
	return sum / len(values)
}

// MaxFloat returns the maximum value in a slice of floats.
func MaxFloat(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	result := values[0]
	for _, num := range values {
		if num > result {
			result = num
		}
	}
	return result
}

// MinFloat returns the minimum value in a slice of floats.
func MinFloat(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	result := values[0]
	for _, num := range values {
		if num < result {
			result = num
		}
	}
	return result
}

// SumFloat returns the sum of a slice of floats.
func SumFloat(values []float64) float64 {
	var result float64
	for _, num := range values {
		result += num
	}
	return result
}

// AverageFloat returns the average of a slice of floats.
func AverageFloat(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var sum float64
	for _, num := range values {
		sum += num
	}
	return sum / float64(len(values))
}

// WeightedAverage returns the weighted average of a slice of floats.
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
