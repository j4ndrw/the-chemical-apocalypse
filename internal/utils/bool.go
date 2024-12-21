package utils

func BoolToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](b bool) T {
	var i T = 0
	if b {
		i = 1
	}
	return i
}
