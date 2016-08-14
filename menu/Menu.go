// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package menu

import (
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/geom"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "Menu.h"
import "C"

var (
	menuMap = make(map[C.platformMenu]*Menu)
	itemMap = make(map[C.platformMenuItem]*Item)
)

// Menu represents a set of menu items.
type Menu struct {
	menu  C.platformMenu
	title string
}

// Bar returns the application menu bar.
func Bar() *Menu {
	if menu, ok := menuMap[C.platformGetMainMenu()]; ok {
		return menu
	}
	menu := NewMenu("")
	C.platformSetMainMenu(menu.menu)
	return menu
}

// NewMenu creates a new Menu.
func NewMenu(title string) *Menu {
	cTitle := C.CString(title)
	menu := &Menu{menu: C.platformNewMenu(cTitle), title: title}
	C.free(unsafe.Pointer(cTitle))
	menuMap[menu.menu] = menu
	return menu
}

// SetServicesMenu marks the specified menu as the services menu.
func SetServicesMenu(menu *Menu) {
	C.platformSetServicesMenu(menu.menu)
}

// SetWindowMenu marks the specified menu as the window menu.
func SetWindowMenu(menu *Menu) {
	C.platformSetWindowMenu(menu.menu)
}

// SetHelpMenu marks the specified menu as the help menu.
func SetHelpMenu(menu *Menu) {
	C.platformSetHelpMenu(menu.menu)
}

// Title returns the title of this Menu.
func (menu *Menu) Title() string {
	return menu.title
}

// Count of Items in this Menu.
func (menu *Menu) Count() int {
	return int(C.platformMenuItemCount(menu.menu))
}

// Item at the specified index, or nil.
func (menu *Menu) Item(index int) *Item {
	if item, ok := itemMap[C.platformGetMenuItem(menu.menu, C.int(index))]; ok {
		return item
	}
	return nil
}

// AddItem creates a new Item and appends it to the end of the Menu.
func (menu *Menu) AddItem(title string, key string) *Item {
	cTitle := C.CString(title)
	cKey := C.CString(key)
	item := &Item{item: C.platformAddMenuItem(menu.menu, cTitle, cKey), title: title}
	C.free(unsafe.Pointer(cTitle))
	C.free(unsafe.Pointer(cKey))
	itemMap[item.item] = item
	return item
}

// AddMenu creates a new sub-Menu and appends it to the end of the Menu.
func (menu *Menu) AddMenu(title string) *Menu {
	item := menu.AddItem(title, "")
	subMenu := NewMenu(title)
	C.platformSetSubMenu(item.item, subMenu.menu)
	return subMenu
}

// AddSeparator creates a new separator and appends it to the end of the Menu.
func (menu *Menu) AddSeparator() {
	item := &Item{item: C.platformAddSeparator(menu.menu)}
	itemMap[item.item] = item
}

// Popup shows the menu at the specified location. If itemAtLocation is specified, it also tries to
// position the menu such that the specified menu item is at that location.
func (menu *Menu) Popup(widget ui.Widget, where geom.Point, itemAtLocation *Item) {
	where = widget.ToWindow(where)
	C.platformPopupMenu(widget.Window().PlatformPtr(), menu.menu, C.float(where.X), C.float(where.Y), itemAtLocation.item)
}

// Dispose of the Menu, releasing any operating system resources it consumed.
func (menu *Menu) Dispose() {
	if menu.menu != nil {
		count := C.platformMenuItemCount(menu.menu)
		var i C.int
		for i = 0; i < count; i++ {
			item := C.platformGetMenuItem(menu.menu, i)
			subMenu := menuMap[C.platformGetSubMenu(item)]
			if subMenu != nil {
				subMenu.Dispose()
			}
			delete(itemMap, item)
		}
		delete(menuMap, menu.menu)
		C.platformDisposeMenu(menu.menu)
		menu.menu = nil
	}
}
