// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package keys

import (
	// #include <X11/keysym.h>
	"C"
)

func init() {
	InsertMapping(' ', &Mapping{KeyCode: Space, KeyChar: ' ', Name: SpaceName})

	InsertMapping(C.XK_Escape, &Mapping{KeyCode: Escape, KeyChar: '\x1b', Name: EscapeName})
	InsertMapping(C.XK_F1, &Mapping{KeyCode: F1, Name: F1Name})
	InsertMapping(C.XK_F2, &Mapping{KeyCode: F2, Name: F2Name})
	InsertMapping(C.XK_F3, &Mapping{KeyCode: F3, Name: F3Name})
	InsertMapping(C.XK_F4, &Mapping{KeyCode: F4, Name: F4Name})
	InsertMapping(C.XK_F5, &Mapping{KeyCode: F5, Name: F5Name})
	InsertMapping(C.XK_F6, &Mapping{KeyCode: F6, Name: F6Name})
	InsertMapping(C.XK_F7, &Mapping{KeyCode: F7, Name: F7Name})
	InsertMapping(C.XK_F8, &Mapping{KeyCode: F8, Name: F8Name})
	InsertMapping(C.XK_F9, &Mapping{KeyCode: F9, Name: F9Name})
	InsertMapping(C.XK_F10, &Mapping{KeyCode: F10, Name: F10Name})
	InsertMapping(C.XK_F11, &Mapping{KeyCode: F11, Name: F11Name})
	InsertMapping(C.XK_F12, &Mapping{KeyCode: F12, Name: F12Name})
	// F13 (aka PrtScn seems to be taken over by the system
	InsertMapping(C.XK_F13, &Mapping{KeyCode: F13, Name: F13Name})
	InsertMapping(C.XK_Scroll_Lock, &Mapping{KeyCode: F14, Name: F14Name})
	InsertMapping(C.XK_Pause, &Mapping{KeyCode: F15, Name: F15Name})

	InsertMapping(C.XK_BackSpace, &Mapping{KeyCode: Backspace, Name: BackspaceName})
	InsertMapping(C.XK_Tab, &Mapping{KeyCode: Tab, KeyChar: '\t', Name: TabName})
	InsertMapping(C.XK_ISO_Left_Tab, &Mapping{KeyCode: Tab, KeyChar: '\t', Name: TabName})
	InsertMapping(C.XK_Caps_Lock, &Mapping{KeyCode: CapsLock, Name: CapsLockName})
	InsertMapping(C.XK_Return, &Mapping{KeyCode: Return, KeyChar: '\n', Name: ReturnName})
	InsertMapping(C.XK_Shift_L, &Mapping{KeyCode: LeftShift, Name: LeftShiftName})
	InsertMapping(C.XK_Shift_R, &Mapping{KeyCode: RightShift, Name: RightShiftName})
	InsertMapping(C.XK_Control_L, &Mapping{KeyCode: LeftControl, Name: LeftControlName})
	InsertMapping(C.XK_Control_R, &Mapping{KeyCode: RightControl, Name: RightControlName})
	InsertMapping(C.XK_Super_L, &Mapping{KeyCode: LeftCmd, Name: LeftCommandName})
	InsertMapping(C.XK_Super_R, &Mapping{KeyCode: RightCmd, Name: RightCommandName})
	InsertMapping(C.XK_Alt_L, &Mapping{KeyCode: LeftOption, Name: LeftAltName})
	InsertMapping(C.XK_Alt_R, &Mapping{KeyCode: RightOption, Name: RightAltName})
	InsertMapping(C.XK_Menu, &Mapping{KeyCode: Menu, Name: MenuName})

	InsertMapping(C.XK_Insert, &Mapping{KeyCode: Insert, Name: InsertName})
	InsertMapping(C.XK_Home, &Mapping{KeyCode: Home, Name: HomeName})
	InsertMapping(C.XK_Page_Up, &Mapping{KeyCode: PageUp, Name: PageUpName})
	InsertMapping(C.XK_Delete, &Mapping{KeyCode: Del, Name: DeleteName})
	InsertMapping(C.XK_End, &Mapping{KeyCode: End, Name: EndName})
	InsertMapping(C.XK_Page_Down, &Mapping{KeyCode: PageDown, Name: PageDownName})

	InsertMapping(C.XK_Up, &Mapping{KeyCode: Up, Name: UpName})
	InsertMapping(C.XK_Left, &Mapping{KeyCode: Left, Name: LeftName})
	InsertMapping(C.XK_Down, &Mapping{KeyCode: Down, Name: DownName})
	InsertMapping(C.XK_Right, &Mapping{KeyCode: Right, Name: RightName})

	InsertMapping(C.XK_Num_Lock, &Mapping{KeyCode: NumLock, Name: NumLockName})
	InsertMapping(C.XK_KP_Divide, &Mapping{KeyCode: NumPadDivide, KeyChar: '/', Name: NumPadDivideName})
	InsertMapping(C.XK_KP_Multiply, &Mapping{KeyCode: NumPadMultiply, KeyChar: '*', Name: NumPadMultiplyName})
	InsertMapping(C.XK_KP_Subtract, &Mapping{KeyCode: NumPadMinus, KeyChar: '-', Name: NumPadMinusName})
	InsertMapping(C.XK_KP_Add, &Mapping{KeyCode: NumPadAdd, KeyChar: '+', Name: NumPadAddName})
	InsertMapping(C.XK_KP_Enter, &Mapping{KeyCode: NumPadEnter, KeyChar: '\n', Name: NumPadEnterName})
	InsertMapping(C.XK_KP_Delete, &Mapping{KeyCode: NumPadDelete, Name: NumPadDeleteName})
	InsertMapping(C.XK_KP_Decimal, &Mapping{KeyCode: NumPadDecimal, KeyChar: '.', Name: NumPadDecimalName})
	InsertMapping(C.XK_KP_1, &Mapping{KeyCode: NumPad1, KeyChar: '1', Name: NumPad1Name})
	InsertMapping(C.XK_KP_2, &Mapping{KeyCode: NumPad2, KeyChar: '2', Name: NumPad2Name})
	InsertMapping(C.XK_KP_3, &Mapping{KeyCode: NumPad3, KeyChar: '3', Name: NumPad3Name})
	InsertMapping(C.XK_KP_4, &Mapping{KeyCode: NumPad4, KeyChar: '4', Name: NumPad4Name})
	InsertMapping(C.XK_KP_5, &Mapping{KeyCode: NumPad5, KeyChar: '5', Name: NumPad5Name})
	InsertMapping(C.XK_KP_6, &Mapping{KeyCode: NumPad6, KeyChar: '6', Name: NumPad6Name})
	InsertMapping(C.XK_KP_7, &Mapping{KeyCode: NumPad7, KeyChar: '7', Name: NumPad7Name})
	InsertMapping(C.XK_KP_8, &Mapping{KeyCode: NumPad8, KeyChar: '8', Name: NumPad8Name})
	InsertMapping(C.XK_KP_9, &Mapping{KeyCode: NumPad9, KeyChar: '9', Name: NumPad9Name})
	InsertMapping(C.XK_KP_0, &Mapping{KeyCode: NumPad0, KeyChar: '0', Name: NumPad0Name})
	InsertMapping(C.XK_KP_Home, &Mapping{KeyCode: NumPadHome, Name: NumPadHomeName})
	InsertMapping(C.XK_KP_Up, &Mapping{KeyCode: NumPadUp, Name: NumPadUpName})
	InsertMapping(C.XK_KP_Page_Up, &Mapping{KeyCode: NumPadPageUp, Name: NumPadPageUpName})
	InsertMapping(C.XK_KP_Left, &Mapping{KeyCode: NumPadLeft, Name: NumPadLeftName})
	InsertMapping(C.XK_KP_Right, &Mapping{KeyCode: NumPadRight, Name: NumPadRightName})
	InsertMapping(C.XK_KP_End, &Mapping{KeyCode: NumPadEnd, Name: NumPadEndName})
	InsertMapping(C.XK_KP_Down, &Mapping{KeyCode: NumPadDown, Name: NumPadDownName})
	InsertMapping(C.XK_KP_Page_Down, &Mapping{KeyCode: NumPadPageDown, Name: NumPadPageDownName})
	InsertMapping(C.XK_KP_Begin, &Mapping{KeyCode: NumPadCenter, Name: NumPadCenterName})
}
