package utils

import "math/rand"

func RandomBetween(start, end int32) int32 {
	if start > end {
		start, end = end, start
	}
	return start + rand.Int31n(end-start+1)
}
