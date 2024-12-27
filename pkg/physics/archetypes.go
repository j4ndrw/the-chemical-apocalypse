package physics

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) Oscillation(amplitude float32, trigValue float64) float32 {
	return amplitude * float32(trigValue)
}
