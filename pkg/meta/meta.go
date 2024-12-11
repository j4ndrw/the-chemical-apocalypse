package meta

type Window struct {
	Width  int32
	Height int32
	Title  string
}

type Meta struct {
	Window    Window
	TargetFPS int32
}

func New(window Window, targetFPS int32) *Meta {
	return &Meta{
		Window: window,
		TargetFPS: targetFPS,
	}
}
