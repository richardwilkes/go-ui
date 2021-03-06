package ui

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/layout"
)

// A Widget is the basic user interface block that interacts with the user.
type Widget interface {
	event.Target

	// Sizer returns the Sizer for this widget, if any.
	Sizer() layout.Sizer
	// SetSizer sets the Sizer for this widget. May be nil.
	SetSizer(sizer layout.Sizer)

	// Layout returns the Layout for this widget, if any.
	Layout() layout.Layout
	// SetLayout sets the Layout for this widget. May be nil.
	SetLayout(layout layout.Layout)
	// NeedLayout returns true if this widget needs to have its children laid
	// out.
	NeedLayout() bool
	// SetNeedLayout sets the whether this widget needs to have its children
	// lait out.
	SetNeedLayout(needLayout bool)
	// LayoutData returns the layout data, if any, associated with this widget.
	LayoutData() interface{}
	// SetLayoutData sets layout data on this widget. May be nil.
	SetLayoutData(data interface{})
	// ValidateLayout triggers any layout that needs to be run by this widget or
	// its children.
	ValidateLayout()

	// Border returns the Border for this widget, if any.
	Border() border.Border
	// SetBorder sets the Border for this widget. May be nil.
	SetBorder(border border.Border)

	// Repaint marks this widget for painting at the next update.
	Repaint()
	// RepaintBounds marks the area 'bounds' in local coordinates within the
	// widget for painting at the next update.
	RepaintBounds(bounds geom.Rect)
	// Paint is called by its owning window when a widget needs to be drawn. 'g'
	// is the graphics context to use. It has already had its clip set to the
	// 'dirty' rectangle. 'dirty' is the area that needs to be drawn.
	Paint(g *draw.Graphics, dirty geom.Rect)

	// Enabled returns true if this widget is currently enabled and can receive
	// events.
	Enabled() bool
	// SetEnabled sets this widget's enabled state.
	SetEnabled(enabled bool)
	// Focusable returns true if this widget can have the keyboard focus.
	Focusable() bool
	// SetFocusable sets whether this widget can have the keyboard focus.
	SetFocusable(focusable bool)
	// Focused returns true if this widget has the keyboard focus.
	Focused() bool
	// GrabFocusWhenClickedOn returns true if this widget wants to have the
	// keyboard focus when it is clicked on.
	GrabFocusWhenClickedOn() bool
	// SetGrabFocusWhenClickedOn sets whether this widget wants to have the
	// keyboard focus when it is clicked on.
	SetGrabFocusWhenClickedOn(grabsFocus bool)

	// Children returns the direct descendents of this widget.
	Children() []Widget
	// IndexOfChild returns the index of the specified child, or -1 if the
	// passed in widget is not a child of this widget.
	IndexOfChild(child Widget) int
	// AddChild adds 'child' as a child of this widget, removing it from any
	// previous parent it may have had.
	AddChild(child Widget)
	// AddChildAtIndex adds 'child' as a child of this widget at the 'index',
	// removing it from any previous parent it may have had. Passing in a
	// negative value for the index will add it to the end.
	AddChildAtIndex(child Widget, index int)
	// RemoveChild removes 'child' from this widget. If 'child' is not a direct
	// descendent of this widget, nothing happens.
	RemoveChild(child Widget)
	// RemoveChildAtIndex removes the child widget at 'index' from this widget.
	// If 'index' is out of range, nothing happens.
	RemoveChildAtIndex(index int)
	// RemoveFromParent removes this widget from its parent, if any.
	RemoveFromParent()
	// Parent returns the parent widget, if any.
	Parent() Widget
	// SetParent sets `parent` to be the parent of this widget. It does not add
	// this widget to the parent as a child. Call AddChild or AddChildAtIndex
	// for that.
	SetParent(parent Widget)
	// Window returns the containing window, if any.
	Window() Window
	// RootOfWindow returns true if this widget is the root widget of a window.
	RootOfWindow() bool

	// Bounds returns the location and size of the widget in its parent's
	// coordinate system.
	Bounds() geom.Rect
	// LocalBounds returns the location and size of the widget in local
	// coordinates.
	LocalBounds() geom.Rect
	// LocalInsetBounds returns the location and size of the widget in local
	// coordinates after adjusting for any Border it may have.
	LocalInsetBounds() geom.Rect
	// SetBounds sets the location and size of the widget in its parent's
	// coordinate system.
	SetBounds(bounds geom.Rect)
	// Location returns the location of this widget in its parent's coordinate
	// system.
	Location() geom.Point
	// SetLocation sets the location of this widget in its parent's coordinate
	// system.
	SetLocation(pt geom.Point)
	// Size returns the size of this widget.
	Size() geom.Size
	// SetSize sets the size of this widget.
	SetSize(size geom.Size)

	// WidgetAt returns the leaf-most child widget containing 'pt', or this
	// widget if no child is found.
	WidgetAt(pt geom.Point) Widget

	// ToWindow converts widget-local coordinates into window coordinates.
	ToWindow(pt geom.Point) geom.Point
	// FromWindow converts window coordinates into widget-local coordinates.
	FromWindow(pt geom.Point) geom.Point

	// Background returns the background color of this widget.
	Background() color.Color
	// SetBackground sets the background color of this widget.
	SetBackground(color color.Color)
}

// Sizes returns the minimum, preferred, and maximum sizes the 'widget' wishes
// to be. It does this by asking the widget's Layout. If no Layout is present,
// then the widget's Sizer is asked. If no Sizer is present, then it finally
// uses a default set of sizes that are used for all components.
func Sizes(widget Widget, hint geom.Size) (min, pref, max geom.Size) {
	if l := widget.Layout(); l != nil {
		return l.Sizes(hint)
	}
	if s := widget.Sizer(); s != nil {
		return s.Sizes(hint)
	}
	return geom.Size{}, geom.Size{}, layout.DefaultMaxSize(geom.Size{})
}
