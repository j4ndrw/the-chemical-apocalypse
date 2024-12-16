package utils

func SliceMap[T, V any](xs []T, fn func(T) V) []V {
	result := make([]V, len(xs))
	for i, t := range xs {
		result[i] = fn(t)
	}
	return result
}
