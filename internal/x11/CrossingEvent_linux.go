// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package x11

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/keys"

	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type CrossingEvent C.XCrossingEvent

func (evt *CrossingEvent) Window() Window {
	return Window(evt.window)
}

func (evt *CrossingEvent) Where() geom.Point {
	return geom.Point{X: float64(evt.x), Y: float64(evt.y)}
}

func (evt *CrossingEvent) Modifiers() keys.Modifiers {
	return Modifiers(evt.state)
}

func (evt *CrossingEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
