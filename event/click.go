package event

import (
	"bytes"
	"fmt"
)

// Click is generated when a widget is being clicked on.
type Click struct {
	target   Target
	finished bool
}

// NewClick creates a new Click event. 'target' is the widget that is being clicked on.
func NewClick(target Target) *Click {
	return &Click{target: target}
}

// Type returns the event type ID.
func (e *Click) Type() Type {
	return ClickType
}

// Target the original target of the event.
func (e *Click) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *Click) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *Click) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *Click) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *Click) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Click[Target: %v", e.target))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
