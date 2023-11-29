package math

import "slices"

type (
	// Number is a generic that can be used in math operations.
	Number interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64
	}

	ErrEmpty  struct{}
	ErrWeight struct{}
	ErrAvg    struct{}
)

func (e ErrEmpty) Error() string {
	return "Numbers cannot be empty"
}

func (e ErrWeight) Error() string {
	return "length of values must equal length of weights"
}

func (e ErrAvg) Error() string {
	return "total weight cannot be zero"
}

// Max returns the maximum value of Numbers.
func Max[T Number](n []T) (T, error) {
	if len(n) == 0 {
		return 0, ErrEmpty{}
	}
	return slices.Max(n), nil
}

// Min returns the minimum value of Numbers.
func Min[T Number](n []T) (T, error) {
	if len(n) == 0 {
		return 0, ErrEmpty{}
	}
	return slices.Min(n), nil
}

// Sum returns the sum of Numbers.
func Sum[T Number](n []T) (T, error) {
	if len(n) == 0 {
		return 0, ErrEmpty{}
	}
	var result T
	for _, num := range n {
		result += num
	}
	return result, nil
}

// Average returns the average of Numbers.
func Average[T Number](n []T) (T, error) {
	if len(n) == 0 {
		return 0, ErrEmpty{}
	}
	sum, err := Sum(n)
	if err != nil {
		return 0, err
	}
	return sum / T(len(n)), nil
}

// WeightedAverage returns the weighted average of Numbers.
func WeightedAverage[T Number](n, weights []T) (T, error) {
	if len(n) == 0 {
		return 0, ErrEmpty{}
	}
	if len(n) != len(weights) {
		return 0, ErrAvg{}
	}
	var sum, totalWeight T
	for i, val := range n {
		sum += val * weights[i]
		totalWeight += weights[i]
	}
	if totalWeight == 0 {
		return 0, ErrAvg{}
	}
	return sum / totalWeight, nil
}
