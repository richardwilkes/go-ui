// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package color

import "strings"

// CSS named colors.
var (
	AliceBlue            = RGB(240, 248, 255)
	AntiqueWhite         = RGB(250, 235, 215)
	Aqua                 = RGB(0, 255, 255)
	Aquamarine           = RGB(127, 255, 212)
	Azure                = RGB(240, 255, 255)
	Beige                = RGB(245, 245, 220)
	Bisque               = RGB(255, 228, 196)
	Black                = RGB(0, 0, 0)
	BlanchedAlmond       = RGB(255, 235, 205)
	Blue                 = RGB(0, 0, 255)
	BlueViolet           = RGB(138, 43, 226)
	Brown                = RGB(165, 42, 42)
	BurlyWood            = RGB(222, 184, 135)
	CadetBlue            = RGB(95, 158, 160)
	Chartreuse           = RGB(127, 255, 0)
	Chocolate            = RGB(210, 105, 30)
	Coral                = RGB(255, 127, 80)
	CornflowerBlue       = RGB(100, 149, 237)
	Cornsilk             = RGB(255, 248, 220)
	Crimson              = RGB(220, 20, 60)
	Cyan                 = RGB(0, 255, 255)
	DarkBlue             = RGB(0, 0, 139)
	DarkCyan             = RGB(0, 139, 139)
	DarkGoldenRod        = RGB(184, 134, 11)
	DarkGray             = RGB(169, 169, 169)
	DarkGreen            = RGB(0, 100, 0)
	DarkGrey             = RGB(169, 169, 169)
	DarkKhaki            = RGB(189, 183, 107)
	DarkMagenta          = RGB(139, 0, 139)
	DarkOliveGreen       = RGB(85, 107, 47)
	DarkOrange           = RGB(255, 140, 0)
	DarkOrchid           = RGB(153, 50, 204)
	DarkRed              = RGB(139, 0, 0)
	DarkSalmon           = RGB(233, 150, 122)
	DarkSeaGreen         = RGB(143, 188, 143)
	DarkSlateBlue        = RGB(72, 61, 139)
	DarkSlateGray        = RGB(47, 79, 79)
	DarkSlateGrey        = RGB(47, 79, 79)
	DarkTurquoise        = RGB(0, 206, 209)
	DarkViolet           = RGB(148, 0, 211)
	DeepPink             = RGB(255, 20, 147)
	DeepSkyBlue          = RGB(0, 191, 255)
	DimGray              = RGB(105, 105, 105)
	DimGrey              = RGB(105, 105, 105)
	DodgerBlue           = RGB(30, 144, 255)
	FireBrick            = RGB(178, 34, 34)
	FloralWhite          = RGB(255, 250, 240)
	ForestGreen          = RGB(34, 139, 34)
	Fuchsia              = RGB(255, 0, 255)
	Gainsboro            = RGB(220, 220, 220)
	GhostWhite           = RGB(248, 248, 255)
	Gold                 = RGB(255, 215, 0)
	GoldenRod            = RGB(218, 165, 32)
	Gray                 = RGB(128, 128, 128)
	Green                = RGB(0, 128, 0)
	GreenYellow          = RGB(173, 255, 47)
	Grey                 = RGB(128, 128, 128)
	HoneyDew             = RGB(240, 255, 240)
	HotPink              = RGB(255, 105, 180)
	IndianRed            = RGB(205, 92, 92)
	Indigo               = RGB(75, 0, 130)
	Ivory                = RGB(255, 255, 240)
	Khaki                = RGB(240, 230, 140)
	Lavender             = RGB(230, 230, 250)
	LavenderBlush        = RGB(255, 240, 245)
	LawnGreen            = RGB(124, 252, 0)
	LemonChiffon         = RGB(255, 250, 205)
	LightBlue            = RGB(173, 216, 230)
	LightCoral           = RGB(240, 128, 128)
	LightCyan            = RGB(224, 255, 255)
	LightGoldenRodYellow = RGB(250, 250, 210)
	LightGray            = RGB(211, 211, 211)
	LightGreen           = RGB(144, 238, 144)
	LightGrey            = RGB(211, 211, 211)
	LightPink            = RGB(255, 182, 193)
	LightSalmon          = RGB(255, 160, 122)
	LightSeaGreen        = RGB(32, 178, 170)
	LightSkyBlue         = RGB(135, 206, 250)
	LightSlateGray       = RGB(119, 136, 153)
	LightSlateGrey       = RGB(119, 136, 153)
	LightSteelBlue       = RGB(176, 196, 222)
	LightYellow          = RGB(255, 255, 224)
	Lime                 = RGB(0, 255, 0)
	LimeGreen            = RGB(50, 205, 50)
	Linen                = RGB(250, 240, 230)
	Magenta              = RGB(255, 0, 255)
	Maroon               = RGB(128, 0, 0)
	MediumAquaMarine     = RGB(102, 205, 170)
	MediumBlue           = RGB(0, 0, 205)
	MediumOrchid         = RGB(186, 85, 211)
	MediumPurple         = RGB(147, 112, 219)
	MediumSeaGreen       = RGB(60, 179, 113)
	MediumSlateBlue      = RGB(123, 104, 238)
	MediumSpringGreen    = RGB(0, 250, 154)
	MediumTurquoise      = RGB(72, 209, 204)
	MediumVioletRed      = RGB(199, 21, 133)
	MidnightBlue         = RGB(25, 25, 112)
	MintCream            = RGB(245, 255, 250)
	MistyRose            = RGB(255, 228, 225)
	Moccasin             = RGB(255, 228, 181)
	NavajoWhite          = RGB(255, 222, 173)
	Navy                 = RGB(0, 0, 128)
	OldLace              = RGB(253, 245, 230)
	Olive                = RGB(128, 128, 0)
	OliveDrab            = RGB(107, 142, 35)
	Orange               = RGB(255, 165, 0)
	OrangeRed            = RGB(255, 69, 0)
	Orchid               = RGB(218, 112, 214)
	PaleGoldenRod        = RGB(238, 232, 170)
	PaleGreen            = RGB(152, 251, 152)
	PaleTurquoise        = RGB(175, 238, 238)
	PaleVioletRed        = RGB(219, 112, 147)
	PapayaWhip           = RGB(255, 239, 213)
	PeachPuff            = RGB(255, 218, 185)
	Peru                 = RGB(205, 133, 63)
	Pink                 = RGB(255, 192, 203)
	Plum                 = RGB(221, 160, 221)
	PowderBlue           = RGB(176, 224, 230)
	Purple               = RGB(128, 0, 128)
	Red                  = RGB(255, 0, 0)
	RosyBrown            = RGB(188, 143, 143)
	RoyalBlue            = RGB(65, 105, 225)
	SaddleBrown          = RGB(139, 69, 19)
	Salmon               = RGB(250, 128, 114)
	SandyBrown           = RGB(244, 164, 96)
	SeaGreen             = RGB(46, 139, 87)
	SeaShell             = RGB(255, 245, 238)
	Sienna               = RGB(160, 82, 45)
	Silver               = RGB(192, 192, 192)
	SkyBlue              = RGB(135, 206, 235)
	SlateBlue            = RGB(106, 90, 205)
	SlateGray            = RGB(112, 128, 144)
	SlateGrey            = RGB(112, 128, 144)
	Snow                 = RGB(255, 250, 250)
	SpringGreen          = RGB(0, 255, 127)
	SteelBlue            = RGB(70, 130, 180)
	Tan                  = RGB(210, 180, 140)
	Teal                 = RGB(0, 128, 128)
	Thistle              = RGB(216, 191, 216)
	Tomato               = RGB(255, 99, 71)
	Turquoise            = RGB(64, 224, 208)
	Violet               = RGB(238, 130, 238)
	Wheat                = RGB(245, 222, 179)
	White                = RGB(255, 255, 255)
	WhiteSmoke           = RGB(245, 245, 245)
	Yellow               = RGB(255, 255, 0)
	YellowGreen          = RGB(154, 205, 50)
	nameToColor          = make(map[string]Color)
	colorToName          = make(map[Color]string)
)

func init() {
	register("AliceBlue", AliceBlue)
	register("AntiqueWhite", AntiqueWhite)
	register("Aqua", Aqua)
	register("Aquamarine", Aquamarine)
	register("Azure", Azure)
	register("Beige", Beige)
	register("Bisque", Bisque)
	register("Black", Black)
	register("BlanchedAlmond", BlanchedAlmond)
	register("Blue", Blue)
	register("BlueViolet", BlueViolet)
	register("Brown", Brown)
	register("BurlyWood", BurlyWood)
	register("CadetBlue", CadetBlue)
	register("Chartreuse", Chartreuse)
	register("Chocolate", Chocolate)
	register("Coral", Coral)
	register("CornflowerBlue", CornflowerBlue)
	register("Cornsilk", Cornsilk)
	register("Crimson", Crimson)
	register("Cyan", Cyan)
	register("DarkBlue", DarkBlue)
	register("DarkCyan", DarkCyan)
	register("DarkGoldenRod", DarkGoldenRod)
	register("DarkGray", DarkGray)
	register("DarkGreen", DarkGreen)
	register("DarkGrey", DarkGrey)
	register("DarkKhaki", DarkKhaki)
	register("DarkMagenta", DarkMagenta)
	register("DarkOliveGreen", DarkOliveGreen)
	register("DarkOrange", DarkOrange)
	register("DarkOrchid", DarkOrchid)
	register("DarkRed", DarkRed)
	register("DarkSalmon", DarkSalmon)
	register("DarkSeaGreen", DarkSeaGreen)
	register("DarkSlateBlue", DarkSlateBlue)
	register("DarkSlateGray", DarkSlateGray)
	register("DarkSlateGrey", DarkSlateGrey)
	register("DarkTurquoise", DarkTurquoise)
	register("DarkViolet", DarkViolet)
	register("DeepPink", DeepPink)
	register("DeepSkyBlue", DeepSkyBlue)
	register("DimGray", DimGray)
	register("DimGrey", DimGrey)
	register("DodgerBlue", DodgerBlue)
	register("FireBrick", FireBrick)
	register("FloralWhite", FloralWhite)
	register("ForestGreen", ForestGreen)
	register("Fuchsia", Fuchsia)
	register("Gainsboro", Gainsboro)
	register("GhostWhite", GhostWhite)
	register("Gold", Gold)
	register("GoldenRod", GoldenRod)
	register("Gray", Gray)
	register("Green", Green)
	register("GreenYellow", GreenYellow)
	register("Grey", Grey)
	register("HoneyDew", HoneyDew)
	register("HotPink", HotPink)
	register("IndianRed", IndianRed)
	register("Indigo", Indigo)
	register("Ivory", Ivory)
	register("Khaki", Khaki)
	register("Lavender", Lavender)
	register("LavenderBlush", LavenderBlush)
	register("LawnGreen", LawnGreen)
	register("LemonChiffon", LemonChiffon)
	register("LightBlue", LightBlue)
	register("LightCoral", LightCoral)
	register("LightCyan", LightCyan)
	register("LightGoldenRodYellow", LightGoldenRodYellow)
	register("LightGray", LightGray)
	register("LightGreen", LightGreen)
	register("LightGrey", LightGrey)
	register("LightPink", LightPink)
	register("LightSalmon", LightSalmon)
	register("LightSeaGreen", LightSeaGreen)
	register("LightSkyBlue", LightSkyBlue)
	register("LightSlateGray", LightSlateGray)
	register("LightSlateGrey", LightSlateGrey)
	register("LightSteelBlue", LightSteelBlue)
	register("LightYellow", LightYellow)
	register("Lime", Lime)
	register("LimeGreen", LimeGreen)
	register("Linen", Linen)
	register("Magenta", Magenta)
	register("Maroon", Maroon)
	register("MediumAquaMarine", MediumAquaMarine)
	register("MediumBlue", MediumBlue)
	register("MediumOrchid", MediumOrchid)
	register("MediumPurple", MediumPurple)
	register("MediumSeaGreen", MediumSeaGreen)
	register("MediumSlateBlue", MediumSlateBlue)
	register("MediumSpringGreen", MediumSpringGreen)
	register("MediumTurquoise", MediumTurquoise)
	register("MediumVioletRed", MediumVioletRed)
	register("MidnightBlue", MidnightBlue)
	register("MintCream", MintCream)
	register("MistyRose", MistyRose)
	register("Moccasin", Moccasin)
	register("NavajoWhite", NavajoWhite)
	register("Navy", Navy)
	register("OldLace", OldLace)
	register("Olive", Olive)
	register("OliveDrab", OliveDrab)
	register("Orange", Orange)
	register("OrangeRed", OrangeRed)
	register("Orchid", Orchid)
	register("PaleGoldenRod", PaleGoldenRod)
	register("PaleGreen", PaleGreen)
	register("PaleTurquoise", PaleTurquoise)
	register("PaleVioletRed", PaleVioletRed)
	register("PapayaWhip", PapayaWhip)
	register("PeachPuff", PeachPuff)
	register("Peru", Peru)
	register("Pink", Pink)
	register("Plum", Plum)
	register("PowderBlue", PowderBlue)
	register("Purple", Purple)
	register("Red", Red)
	register("RosyBrown", RosyBrown)
	register("RoyalBlue", RoyalBlue)
	register("SaddleBrown", SaddleBrown)
	register("Salmon", Salmon)
	register("SandyBrown", SandyBrown)
	register("SeaGreen", SeaGreen)
	register("SeaShell", SeaShell)
	register("Sienna", Sienna)
	register("Silver", Silver)
	register("SkyBlue", SkyBlue)
	register("SlateBlue", SlateBlue)
	register("SlateGray", SlateGray)
	register("SlateGrey", SlateGrey)
	register("Snow", Snow)
	register("SpringGreen", SpringGreen)
	register("SteelBlue", SteelBlue)
	register("Tan", Tan)
	register("Teal", Teal)
	register("Thistle", Thistle)
	register("Tomato", Tomato)
	register("Turquoise", Turquoise)
	register("Violet", Violet)
	register("Wheat", Wheat)
	register("White", White)
	register("WhiteSmoke", WhiteSmoke)
	register("Yellow", Yellow)
	register("YellowGreen", YellowGreen)
}

func register(name string, color Color) {
	nameToColor[strings.ToLower(name)] = color
	colorToName[color] = name
}
