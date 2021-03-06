package separator

import (
	"fmt"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/widget"
)

// Separator provides a simple vertical or horizontal separator line.
type Separator struct {
	widget.Block
	horizontal bool
}

// New creates a new separator.
func New(horizontal bool) *Separator {
	sep := &Separator{}
	sep.InitTypeAndID(sep)
	sep.Initialize(horizontal)
	return sep
}

// Initialize the separator.
func (sep *Separator) Initialize(horizontal bool) {
	sep.horizontal = horizontal
	sep.Describer = func() string { return fmt.Sprintf("Separator #%d", sep.ID()) }
	sep.SetSizer(sep)
	sep.EventHandlers().Add(event.PaintType, sep.paint)
}

// Sizes implements Sizer
func (sep *Separator) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	if sep.horizontal {
		if hint.Width == layout.NoHint {
			pref.Width = 1
		} else {
			pref.Width = hint.Width
		}
		min.Width = 1
		max.Width = layout.DefaultMax
		min.Height = 1
		pref.Height = 1
		max.Height = 1
	} else {
		if hint.Height == layout.NoHint {
			pref.Height = 1
		} else {
			pref.Height = hint.Height
		}
		min.Height = 1
		max.Height = layout.DefaultMax
		min.Width = 1
		pref.Width = 1
		max.Width = 1
	}
	if border := sep.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	return min, pref, max
}

func (sep *Separator) paint(evt event.Event) {
	bounds := sep.LocalInsetBounds()
	if sep.horizontal {
		if bounds.Height > 1 {
			bounds.Y += (bounds.Height - 1) / 2
			bounds.Height = 1
		}
	} else {
		if bounds.Width > 1 {
			bounds.X += (bounds.Width - 1) / 2
			bounds.Width = 1
		}
	}
	gc := evt.(*event.Paint).GC()
	gc.SetColor(color.Background.AdjustBrightness(-0.25))
	gc.FillRect(bounds)
}
