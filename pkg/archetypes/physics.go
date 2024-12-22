package archetypes

type physics struct{}

var Physics = physics{}

func (_ *physics) Oscillation(amplitude float32, trigValue float64) float32 {
	return amplitude * float32(trigValue)
}
