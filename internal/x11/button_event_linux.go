package x11

import (
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/event/button"
	"github.com/richardwilkes/ui/keys"
)

type ButtonEvent C.XButtonEvent

func (evt *ButtonEvent) Window() Window {
	return Window(evt.window)
}

func (evt *ButtonEvent) Where() geom.Point {
	return geom.Point{X: float64(evt.x), Y: float64(evt.y)}
}

func (evt *ButtonEvent) Modifiers() keys.Modifiers {
	return Modifiers(evt.state)
}

func (evt *ButtonEvent) Button() int {
	switch evt.button {
	case 1:
		return button.Left
	case 2:
		return button.Middle
	case 3:
		return button.Right
	default:
		return -1
	}
}

func (evt *ButtonEvent) IsScrollWheel() bool {
	return evt.button > 3 && evt.button < 8
}

func (evt *ButtonEvent) ScrollWheelDirection() geom.Point {
	var result geom.Point
	switch evt.button {
	case 4: // Up
		result.Y = -1
	case 5: // Down
		result.Y = 1
	case 6: // Left
		result.X = -1
	case 7: // Right
		result.X = 1
	}
	return result
}

func (evt *ButtonEvent) When() C.Time {
	return evt.time
}

func (evt *ButtonEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
