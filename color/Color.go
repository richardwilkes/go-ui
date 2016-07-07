// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package color

import (
	"fmt"
	"strconv"
	"strings"
)

// Color contains the value of a color used for drawing, stored as 0xAARRGGBB.
type Color uint32

// String -- implements the fmt.Stringer interface. The output can be used as a color in CSS.
func (c Color) String() string {
	if name, ok := colorToName[c]; ok {
		return name
	}
	if c.HasAlpha() {
		return fmt.Sprintf("rgba(%d,%d,%d,%v)", c.Red(), c.Green(), c.Blue(), c.AlphaIntensity())
	}
	return fmt.Sprintf("#%06X", uint32(c&0xFFFFFF))
}

// Decode creates a Color from a string. The string may be in any of the standard CSS formats:
//
// - CSS predefined color name, e.g. "Yellow"
// - CSS rgb(), e.g. "rgb(255, 255, 0)"
// - CSS rgba(), e.g. "rgba(255, 255, 0, 0.3)"
// - CSS short hexadecimal colors, e.g. "#FF0"
// - CSS long hexadecimal colors, e.g. "#FFFF00"
// - CCS hsl(), e.g. "hsl(120, 100%, 50%)"
// - CSS hsla(), e.g. "hsla(120, 100%, 50%, 0.3)"
func Decode(buffer string) Color {
	buffer = strings.ToLower(strings.TrimSpace(buffer))
	if color, ok := nameToColor[buffer]; ok {
		return color
	}
	switch {
	case strings.HasPrefix(buffer, "#"):
		buffer = buffer[1:]
		switch len(buffer) {
		case 3:
			return RGB(extractChannel(buffer[0:1]+buffer[0:1], 16), extractChannel(buffer[1:2]+buffer[1:2], 16), extractChannel(buffer[2:3]+buffer[2:3], 16))
		case 6:
			return RGB(extractChannel(buffer[0:2], 16), extractChannel(buffer[2:4], 16), extractChannel(buffer[4:6], 16))
		}
	case strings.HasPrefix(buffer, "rgb(") && strings.HasSuffix(buffer, ")"):
		parts := strings.SplitN(strings.TrimSpace(buffer[4:len(buffer)-1]), ",", 4)
		if len(parts) == 3 {
			return RGB(extractChannel(parts[0], 10), extractChannel(parts[1], 10), extractChannel(parts[2], 10))
		}
	case strings.HasPrefix(buffer, "rgba(") && strings.HasSuffix(buffer, ")"):
		parts := strings.SplitN(strings.TrimSpace(buffer[5:len(buffer)-1]), ",", 5)
		if len(parts) == 4 {
			return RGBA(extractChannel(parts[0], 10), extractChannel(parts[1], 10), extractChannel(parts[2], 10), extractAlpha(parts[3]))
		}
	case strings.HasPrefix(buffer, "hsl(") && strings.HasSuffix(buffer, ")"):
		parts := strings.SplitN(strings.TrimSpace(buffer[4:len(buffer)-1]), ",", 4)
		if len(parts) == 3 {
			return HSB(float32(extractChannel(parts[0], 10))/360, extractPercentage(parts[1]), extractPercentage(parts[2]))
		}
	case strings.HasPrefix(buffer, "hsla(") && strings.HasSuffix(buffer, ")"):
		parts := strings.SplitN(strings.TrimSpace(buffer[5:len(buffer)-1]), ",", 5)
		if len(parts) == 4 {
			return HSBA(float32(extractChannel(parts[0], 10))/360, extractPercentage(parts[1]), extractPercentage(parts[2]), extractAlpha(parts[3]))
		}
	}
	return 0
}

func extractChannel(buffer string, base int) int {
	if value, err := strconv.ParseInt(strings.TrimSpace(buffer), base, 64); err == nil {
		return int(value)
	}
	return 0
}

func extractAlpha(buffer string) float32 {
	alpha, err := strconv.ParseFloat(strings.TrimSpace(buffer), 32)
	if err != nil {
		return 0
	}
	return clamp0To1(float32(alpha))
}

func extractPercentage(buffer string) float32 {
	buffer = strings.TrimSpace(buffer)
	if strings.HasSuffix(buffer, "%") {
		value, err := strconv.Atoi(strings.TrimSpace(buffer[:len(buffer)-1]))
		if err != nil {
			return 0
		}
		return clamp0To1(float32(value) / 100)
	}
	return 0
}

// HasAlpha returns true if the color is not fully opaque.
func (c Color) HasAlpha() bool {
	return (c & 0xFF000000) != 0xFF000000
}

// Alpha returns the alpha channel, in the range of 0-255.
func (c Color) Alpha() int {
	return int((c >> 24) & 0xFF)
}

// SetAlpha returns a new color based on this color, but with the alpha channel replaced.
func (c Color) SetAlpha(alpha int) Color {
	return Color((clamp0To255(alpha) << 24) | (int(c) & 0x00FFFFFF))
}

// AlphaIntensity returns the alpha channel, in the range of 0-1.
func (c Color) AlphaIntensity() float32 {
	return float32(c.Alpha()) / 255
}

// SetAlphaIntensity returns a new color based on this color, but with the alpha channel replaced.
func (c Color) SetAlphaIntensity(alpha float32) Color {
	return RGBA(c.Red(), c.Green(), c.Blue(), alpha)
}

// Monochrome returns true if each color component is the same, making it a shade of gray.
func (c Color) Monochrome() bool {
	green := c.Green()
	return c.Red() == green && green == c.Blue()
}

// Luminance returns a value from 0-1 representing the perceived brightness.
// Lower values represent darker colors, while higher values represent brighter colors.
func (c Color) Luminance() float32 {
	return 0.299*c.RedIntensity() + 0.587*c.GreenIntensity() + 0.114*c.BlueIntensity()
}

// Blend blends this color with another color. pct is the amount of the other color to use.
func (c Color) Blend(other Color, pct float32) Color {
	pct = clamp0To1(pct)
	rem := 1 - pct
	return RGBA(scaleAndClamp(c.RedIntensity()*rem+other.RedIntensity()*pct), scaleAndClamp(c.GreenIntensity()*rem+other.GreenIntensity()*pct), scaleAndClamp(c.BlueIntensity()*rem+other.BlueIntensity()*pct), c.AlphaIntensity())
}