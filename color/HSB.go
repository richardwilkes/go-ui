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
	"math"
)

// HSB creates a new opaque Color from HSB values in the range 0-1.
func HSB(hue, saturation, brightness float32) Color {
	return HSBA(hue, saturation, brightness, 1)
}

// HSBA creates a new Color from HSBA values in the range 0-1.
func HSBA(hue, saturation, brightness, alpha float32) Color {
	saturation = clamp0To1(saturation)
	brightness = clamp0To1(brightness)
	v := scaleAndClamp(brightness)
	if saturation == 0 {
		return RGBA(v, v, v, alpha)
	}
	h := (hue - float32(math.Floor(float64(hue)))) * 6
	f := h - float32(math.Floor(float64(h)))
	p := scaleAndClamp(brightness * (1 - saturation))
	q := scaleAndClamp(brightness * (1 - saturation*f))
	t := scaleAndClamp(brightness * (1 - (saturation * (1 - f))))
	switch int(h) {
	case 0:
		return RGBA(v, t, p, alpha)
	case 1:
		return RGBA(q, v, p, alpha)
	case 2:
		return RGBA(p, v, t, alpha)
	case 3:
		return RGBA(p, q, v, alpha)
	case 4:
		return RGBA(t, p, v, alpha)
	default:
		return RGBA(v, p, q, alpha)
	}
}

// Hue of the color, a value from 0-1.
func (c Color) Hue() float32 {
	hue, _, _ := c.HSB()
	return hue
}

// SetHue creates a new color from this color with the specified hue, a value from 0-1.
func (c Color) SetHue(hue float32) Color {
	_, s, b := c.HSB()
	return HSBA(hue, s, b, c.AlphaIntensity())
}

// AdjustHue creates a new color from this color with its hue adjusted by the specified amount.
func (c Color) AdjustHue(amount float32) Color {
	h, s, b := c.HSB()
	return HSBA(h+amount, s, b, c.AlphaIntensity())
}

// Saturation of the color, a value from 0-1.
func (c Color) Saturation() float32 {
	brightness := c.Brightness()
	if brightness != 0 {
		return (brightness - (float32(minOf3(c.Red(), c.Green(), c.Blue())) / 255)) / brightness
	}
	return 0
}

// SetSaturation creates a new color from this color with the specified saturation.
func (c Color) SetSaturation(saturation float32) Color {
	h, _, b := c.HSB()
	return HSBA(h, saturation, b, c.AlphaIntensity())
}

// AdjustSaturation creates a new color from this color with its saturation adjusted by the
// specified amount.
func (c Color) AdjustSaturation(amount float32) Color {
	h, s, b := c.HSB()
	return HSBA(h, s+amount, b, c.AlphaIntensity())
}

// Brightness of the color, a value from 0-1.
func (c Color) Brightness() float32 {
	return float32(maxOf3(c.Red(), c.Green(), c.Blue())) / 255
}

// SetBrightness creates a new color from this color with the specified brightness.
func (c Color) SetBrightness(brightness float32) Color {
	h, s, _ := c.HSB()
	return HSBA(h, s, brightness, c.AlphaIntensity())
}

// AdjustBrightness creates a new color from this color with its brightness adjusted by the
// specified amount.
func (c Color) AdjustBrightness(amount float32) Color {
	h, s, b := c.HSB()
	return HSBA(h, s, b+amount, c.AlphaIntensity())
}

// HSB returns the hue, saturation and brightness of the color. Values are in the range 0-1.
func (c Color) HSB() (hue, saturation, brightness float32) {
	red := c.Red()
	green := c.Green()
	blue := c.Blue()
	cmax := maxOf3(red, green, blue)
	cmin := minOf3(red, green, blue)
	if cmax != 0 {
		saturation = float32(cmax-cmin) / float32(cmax)
	} else {
		saturation = 0
	}
	if saturation == 0 {
		hue = 0
	} else {
		div := float32(cmax - cmin)
		r := float32(cmax-red) / div
		g := float32(cmax-green) / div
		b := float32(cmax-blue) / div
		if r == float32(cmax) {
			hue = b - g
		} else if g == float32(cmax) {
			hue = 2 + r - b
		} else {
			hue = 4 + g - r
		}
		hue /= 6
		if hue < 0 {
			hue++
		}
	}
	return hue, saturation, float32(cmax) / 255
}

func minOf3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func maxOf3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func clamp0To1(value float32) float32 {
	switch {
	case value < 0:
		return 0
	case value > 1:
		return 1
	default:
		return value
	}
}

func clamp0To255(value int) int {
	switch {
	case value < 0:
		return 0
	case value > 255:
		return 255
	default:
		return value
	}
}

func scaleAndClamp(value float32) int {
	return clamp0To255(int(clamp0To1(value)*255 + 0.5))
}
