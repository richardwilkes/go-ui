package event

import (
	"bytes"
	"fmt"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/draw"
)

// Paint is generated when a widget needs to be drawn.
type Paint struct {
	target   Target
	gc       *draw.Graphics
	dirty    geom.Rect
	finished bool
}

// NewPaint creates a new Paint event. 'target' is the widget to be painted. 'gc' is a graphics
// context setup for drawing into the widget. 'dirty' is the area within the widget that needs to
// be redrawn.
func NewPaint(target Target, gc *draw.Graphics, dirty geom.Rect) *Paint {
	return &Paint{target: target, gc: gc, dirty: dirty}
}

// Type returns the event type ID.
func (e *Paint) Type() Type {
	return PaintType
}

// Target the original target of the event.
func (e *Paint) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *Paint) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *Paint) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *Paint) Finish() {
	e.finished = true
}

// GC returns the graphics context to use when drawing the widget.
func (e *Paint) GC() *draw.Graphics {
	return e.gc
}

// DirtyRect returns the area within the widget that needs to be redrawn.
func (e *Paint) DirtyRect() geom.Rect {
	return e.dirty
}

// String implements the fmt.Stringer interface.
func (e *Paint) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Paint[Dirty: [%v], Target: %v", e.dirty, e.target))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
