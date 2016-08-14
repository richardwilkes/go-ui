// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package main

import (
	"fmt"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/Demo/images"
	"github.com/richardwilkes/ui/app"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/font"
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget"
	"unicode"
)

var (
	aboutWindow ui.Window
)

func main() {
	// event.TraceAllEvents = true
	// event.TraceEventTypes = append(event.TraceEventTypes, event.MouseDownType, event.MouseDraggedType, event.MouseUpType)
	app.App.EventHandlers().Add(event.AppWillFinishStartupType, func(evt event.Event) {
		createMenuBar()
		createButtonsWindow()
	})
	app.Start()
}

func createMenuBar() {
	_, aboutItem, prefsItem := widget.AddAppMenu()
	aboutItem.EventHandlers().Add(event.SelectionType, createAboutWindow)
	prefsItem.EventHandlers().Add(event.SelectionType, createPreferencesWindow)

	fileMenu := menu.Bar().AddMenu("File")
	fileMenu.AddItem("Open", "o")

	m := menu.Bar().AddMenu("Edit")
	widget.AddCutItem(m)
	widget.AddCopyItem(m)
	widget.AddPasteItem(m)
	m.AddSeparator()
	widget.AddDeleteItem(m)
	widget.AddSelectAllItem(m)

	widget.AddWindowMenu()
	widget.AddHelpMenu()
}

func createButtonsWindow() {
	wnd := widget.NewWindow(geom.Point{}, widget.StdWindowMask)
	wnd.SetTitle("Demo")

	root := wnd.RootWidget()
	root.SetBorder(border.NewEmpty(geom.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10}))
	layout.NewPrecision(root).SetVerticalSpacing(10)

	buttonsPanel := createButtonsPanel()
	buttonsPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(buttonsPanel)

	addSeparator(root)

	checkBoxPanel := createCheckBoxPanel()
	checkBoxPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(checkBoxPanel)

	addSeparator(root)

	radioButtonsPanel := createRadioButtonsPanel()
	radioButtonsPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(radioButtonsPanel)

	addSeparator(root)

	popupMenusPanel := createPopupMenusPanel()
	popupMenusPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(popupMenusPanel)

	addSeparator(root)

	wrapper := widget.NewBlock()
	layout.NewPrecision(wrapper).SetColumns(2).SetEqualColumns(true).SetHorizontalSpacing(10)
	wrapper.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	textFieldsPanel := createTextFieldsPanel()
	textFieldsPanel.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	wrapper.AddChild(textFieldsPanel)
	wrapper.AddChild(createListPanel())
	root.AddChild(wrapper)

	addSeparator(root)

	img, err := draw.AcquireImageFromURL("http://allwallpapersnew.com/wp-content/gallery/stock-photos-for-free/grassy_field_sunset___free_stock_by_kevron2001-d5blgkr.jpg")
	if err == nil {
		content := widget.NewImageLabel(img)
		content.SetFocusable(true)
		_, prefSize, _ := layout.Sizes(content, layout.NoHintSize)
		content.SetSize(prefSize)
		scrollArea := widget.NewScrollArea(content, widget.ScrollContentUnmodified)
		scrollArea.SetLayoutData(layout.NewPrecisionData().SetHorizontalAlignment(draw.AlignFill).SetVerticalAlignment(draw.AlignFill).SetHorizontalGrab(true).SetVerticalGrab(true))
		root.AddChild(scrollArea)
	} else {
		fmt.Println(err)
	}

	wnd.Pack()
	wnd.ToFront()
}

func createListPanel() ui.Widget {
	list := widget.NewList(&widget.TextCellFactory{})
	list.Append("One",
		"Two",
		"Three with some long text to make it interesting",
		"Four",
		"Five")
	list.EventHandlers().Add(event.SelectionType, func(evt event.Event) {
		fmt.Print("Selection changed in list. Now:")
		index := -1
		first := true
		for {
			index = list.Selection.NextSet(index + 1)
			if index == -1 {
				break
			}
			if first {
				first = false
			} else {
				fmt.Print(",")
			}
			fmt.Printf(" %d", index)
		}
		fmt.Println()
	})
	list.EventHandlers().Add(event.ClickType, func(evt event.Event) {
		fmt.Println("Double-clicked on list")
	})
	_, prefSize, _ := layout.Sizes(list, layout.NoHintSize)
	list.SetSize(prefSize)
	scrollArea := widget.NewScrollArea(list, widget.ScrollContentFill)
	scrollArea.SetLayoutData(layout.NewPrecisionData().SetHorizontalAlignment(draw.AlignFill).SetVerticalAlignment(draw.AlignFill).SetHorizontalGrab(true).SetVerticalGrab(true))
	return scrollArea
}

func addSeparator(root ui.Widget) {
	sep := widget.NewSeparator(true)
	sep.SetLayoutData(layout.NewPrecisionData().SetHorizontalAlignment(draw.AlignFill))
	root.AddChild(sep)
}

func createButtonsPanel() ui.Widget {
	panel := widget.NewBlock()
	layout.NewFlow(panel).SetHorizontalSpacing(5).SetVerticalSpacing(5).SetVerticallyCentered(true)

	createButton("Press Me", panel)
	createButton("Disabled", panel).SetEnabled(false)

	img, err := draw.AcquireImageFromFile(images.FS, "/home.png")
	if err == nil {
		createImageButton(img, "Home", panel)
		createImageButton(img, "Home (disabled)", panel).SetEnabled(false)
	} else {
		fmt.Println(err)
	}

	img, err = draw.AcquireImageFromFile(images.FS, "/classic-apple-logo.png")
	if err == nil {
		createImageButton(img, "Classic Apple Logo", panel)
		createImageButton(img, "Classic Apple Logo (disabled)", panel).SetEnabled(false)
	} else {
		fmt.Println(err)
	}

	return panel
}

func createButton(title string, panel ui.Widget) *widget.Button {
	button := widget.NewButton(title)
	button.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The button '%s' was clicked.\n", title) })
	widget.NewSimpleToolTip(button, fmt.Sprintf("This is the tooltip for the '%s' button.", title))
	panel.AddChild(button)
	return button
}

func createImageButton(img *draw.Image, name string, panel ui.Widget) *widget.ImageButton {
	size := img.Size()
	size.Width /= 2
	size.Height /= 2
	button := widget.NewImageButtonWithImageSize(img, size)
	button.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The button '%s' was clicked.\n", name) })
	widget.NewSimpleToolTip(button, name)
	panel.AddChild(button)
	return button
}

func createCheckBoxPanel() ui.Widget {
	panel := widget.NewBlock()
	layout.NewPrecision(panel)
	createCheckBox("Press Me", panel)
	createCheckBox("Initially Mixed", panel).SetState(widget.Mixed)
	createCheckBox("Disabled", panel).SetEnabled(false)
	checkbox := createCheckBox("Disabled w/Check", panel)
	checkbox.SetEnabled(false)
	checkbox.SetState(widget.Checked)
	return panel
}

func createCheckBox(title string, panel ui.Widget) *widget.CheckBox {
	checkbox := widget.NewCheckBox(title)
	checkbox.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The checkbox '%s' was clicked.\n", title) })
	widget.NewSimpleToolTip(checkbox, fmt.Sprintf("This is the tooltip for the '%s' checkbox.", title))
	panel.AddChild(checkbox)
	return checkbox
}

func createRadioButtonsPanel() ui.Widget {
	panel := widget.NewBlock()
	layout.NewPrecision(panel)

	group := widget.NewRadioButtonGroup()
	first := createRadioButton("First", panel, group)
	createRadioButton("Second", panel, group)
	createRadioButton("Third (disabled)", panel, group).SetEnabled(false)
	createRadioButton("Fourth", panel, group)
	group.Select(first)

	return panel
}

func createRadioButton(title string, panel ui.Widget, group *widget.RadioButtonGroup) *widget.RadioButton {
	rb := widget.NewRadioButton(title)
	rb.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The radio button '%s' was clicked.\n", title) })
	widget.NewSimpleToolTip(rb, fmt.Sprintf("This is the tooltip for the '%s' radio button.", title))
	panel.AddChild(rb)
	group.Add(rb)
	return rb
}

func createPopupMenusPanel() ui.Widget {
	panel := widget.NewBlock()
	layout.NewPrecision(panel)

	createPopupMenu(panel, 1, "One", "Two", "Three", "", "Four", "Five", "Six")
	createPopupMenu(panel, 2, "Red", "Blue", "Green").SetEnabled(false)

	return panel
}

func createPopupMenu(panel ui.Widget, selection int, titles ...string) *widget.PopupMenu {
	p := widget.NewPopupMenu()
	widget.NewSimpleToolTip(p, fmt.Sprintf("This is the tooltip for the PopupMenu with %d items.", len(titles)))
	for _, title := range titles {
		if title == "" {
			p.AddSeparator()
		} else {
			p.AddItem(title)
		}
	}
	p.SelectIndex(selection)
	p.EventHandlers().Add(event.SelectionType, func(evt event.Event) { fmt.Printf("The '%v' item was selected from the PopupMenu.\n", p.Selected()) })
	panel.AddChild(p)
	return p
}

func createTextFieldsPanel() ui.Widget {
	panel := widget.NewBlock()
	layout.NewPrecision(panel)

	createTextField("First Text Field", panel)
	createTextField("Second Text Field (disabled)", panel).SetEnabled(false)
	createTextField("", panel).SetWatermark("Watermarked")
	field := createTextField("", panel)
	field.SetWatermark("Enter only numbers")
	field.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		e := evt.(*event.Validate)
		for _, r := range field.Text() {
			if !unicode.IsDigit(r) {
				e.MarkInvalid()
				break
			}
		}
	})

	return panel
}

func createTextField(text string, panel ui.Widget) *widget.TextField {
	field := widget.NewTextField()
	field.SetText(text)
	field.SetLayoutData(layout.NewPrecisionData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	widget.NewSimpleToolTip(field, fmt.Sprintf("This is the tooltip for the '%s' text field.", text))
	panel.AddChild(field)
	return field
}

func createAboutWindow(evt event.Event) {
	if aboutWindow == nil {
		aboutWindow = widget.NewWindow(geom.Point{}, widget.TitledWindowMask|widget.ClosableWindowMask)
		aboutWindow.EventHandlers().Add(event.ClosedType, func(evt event.Event) { aboutWindow = nil })
		aboutWindow.SetTitle("About " + app.Name())
		root := aboutWindow.RootWidget()
		root.SetBorder(border.NewEmpty(geom.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10}))
		layout.NewPrecision(root)
		title := widget.NewLabelWithFont(app.Name(), font.Acquire(font.EmphasizedSystemDesc))
		title.SetLayoutData(layout.NewPrecisionData().SetHorizontalAlignment(draw.AlignMiddle))
		root.AddChild(title)
		desc := widget.NewLabel("Simple app to demonstrate the\ncapabilities of the ui framework.")
		root.AddChild(desc)
		aboutWindow.Pack()
	}
	aboutWindow.ToFront()
}

func createPreferencesWindow(evt event.Event) {
	fmt.Println("Preferences...")
}

type scrollTarget struct {
	hpos float32
	vpos float32
}

// LineScrollAmount implements ui.Scrollable.
func (st *scrollTarget) LineScrollAmount(horizontal, towardsStart bool) float32 {
	return 1
}

// PageScrollAmount implements ui.Scrollable.
func (st *scrollTarget) PageScrollAmount(horizontal, towardsStart bool) float32 {
	return 10
}

// ScrolledPosition implements ui.Scrollable.
func (st *scrollTarget) ScrolledPosition(horizontal bool) float32 {
	if horizontal {
		return st.hpos
	}
	return st.vpos
}

// SetScrolledPosition implements ui.Scrollable.
func (st *scrollTarget) SetScrolledPosition(horizontal bool, position float32) {
	if horizontal {
		st.hpos = position
	} else {
		st.vpos = position
	}
}

// VisibleSize implements ui.Scrollable.
func (st *scrollTarget) VisibleSize(horizontal bool) float32 {
	return 10
}

// ContentSize implements ui.Scrollable.
func (st *scrollTarget) ContentSize(horizontal bool) float32 {
	return 1000
}
