// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package menu

func platformMenuBar() PlatformMenu {
	// RAW: Implement for Windows
	return nil
}

func platformNewMenu(title string) PlatformMenu {
	// RAW: Implement for Windows
	return nil
}

func (menu *Menu) platformItem(index int) PlatformItem {
	// RAW: Implement for Windows
	return nil
}

func (menu *Menu) platformAddItem(title string, key string) PlatformItem {
	// RAW: Implement for Windows
	return nil
}

func (menu *Menu) platformAddSeparator() PlatformItem {
	// RAW: Implement for Windows
	return nil
}

func (menu *Menu) platformCount() int {
	// RAW: Implement for Windows
	return 0
}

func (menu *Menu) platformSetAsMenuBar() {
	// RAW: Implement for Windows
}

func (menu *Menu) platformSetAsServicesMenu() {
	// RAW: Implement for Windows
}

func (menu *Menu) platformSetAsWindowMenu() {
	// RAW: Implement for Windows
}

func (menu *Menu) platformSetAsHelpMenu() {
	// RAW: Implement for Windows
}

func (menu *Menu) platformDispose() {
	// RAW: Implement for Windows
}
