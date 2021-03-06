package event

import (
	"bytes"
	"fmt"
)

// Closed is generated when a window is closed.
type Closed struct {
	target   Target
	finished bool
}

// NewClosed creates a new Closed event. 'target' is the window that was closed.
func NewClosed(target Target) *Closed {
	return &Closed{target: target}
}

// Type returns the event type ID.
func (e *Closed) Type() Type {
	return ClosedType
}

// Target the original target of the event.
func (e *Closed) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *Closed) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *Closed) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *Closed) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *Closed) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Closed[Target: %v", e.target))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
