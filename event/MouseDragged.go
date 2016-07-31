// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

import (
	"bytes"
	"fmt"
	"github.com/richardwilkes/ui/geom"
	"reflect"
)

// MouseDragged is generated when the mouse is moved within a widget while a mouse button
// is down.
type MouseDragged struct {
	target    Target
	where     geom.Point
	modifiers KeyMask
	finished  bool
	discarded bool
}

// NewMouseDragged creates a new MouseDragged event. 'target' is the widget that was being clicked
// on. 'where' is the location in the window where the mouse is. 'modifiers' are the keyboard
// modifiers keys that were down.
func NewMouseDragged(target Target, where geom.Point, modifiers KeyMask) *MouseDragged {
	return &MouseDragged{target: target, where: where, modifiers: modifiers}
}

// Type returns the event type ID.
func (e *MouseDragged) Type() Type {
	return MouseDraggedType
}

// Target the original target of the event.
func (e *MouseDragged) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *MouseDragged) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *MouseDragged) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *MouseDragged) Finish() {
	e.finished = true
}

// Where returns the location in the window the mouse is being pressed.
func (e *MouseDragged) Where() geom.Point {
	return e.where
}

// Modifiers returns the key modifiers that were down.
func (e *MouseDragged) Modifiers() KeyMask {
	return e.modifiers
}

// String implements the fmt.Stringer interface.
func (e *MouseDragged) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("MouseDragged[Where: [%v], Target: %v", e.where, reflect.ValueOf(e.target).Pointer()))
	modifiers := e.modifiers.String()
	if modifiers != "" {
		buffer.WriteString(", ")
		buffer.WriteString(modifiers)
	}
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}