package common

import "golang.org/x/exp/constraints"

// Sum sums all the values in the slice
func Sum[T int | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// Mul multiplies all the values in the slice
func Mul[T int | float64](s []T) T {
	var mul T = 1
	for _, v := range s {
		mul *= v
	}
	return mul
}

// Abs returns the absolute value of the input
func Abs[T int | float64](s T) T {
	if s < 0 {
		return s * -1
	}
	return s
}

// Max returns the maximum value in the slice
func Max[T constraints.Ordered](s ...T) T {
	var max T
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

// Min returns the minimum value in the slice
func Min[T constraints.Ordered](s ...T) T {
	var min T
	for i, v := range s {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}
