package common

// Sum sums all the values in the slice
func Sum[T int | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}