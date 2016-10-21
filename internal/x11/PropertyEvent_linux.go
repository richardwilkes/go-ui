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
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type PropertyEvent C.XPropertyEvent

func (evt *PropertyEvent) Time() C.Time {
	return evt.time
}

func (evt *PropertyEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
