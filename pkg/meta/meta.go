package meta

type Window struct {
	Width  int32
	Height int32
	Title  string
	Screen struct {
		Windowed   bool
		Fullscreen bool
		Borderless bool
	}
}

type Meta struct {
	Window    Window
	TargetFPS int32
}

func New(window Window, targetFPS int32) *Meta {
	return &Meta{
		Window:    window,
		TargetFPS: targetFPS,
	}
}

func Default() *Meta {
	return New(Window{
		Title:      "The Chemical Apocalypse",
		Width:      1280,
		Height:     720,
	}, 60)
}
