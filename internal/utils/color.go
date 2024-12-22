package utils

func ColorOpacity(percentage float32) uint8 {
	Assert(
		0 <= percentage && percentage <= 1,
		"Color opacity must be a float value between 0 and 1",
	)
	return uint8(percentage * 255)
}
