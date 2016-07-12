// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// ImageLabel represents a non-interactive image.
type ImageLabel struct {
	Block
	image *Image
}

// NewImageLabel creates an ImageLabel with the specified image.
func NewImageLabel(img *Image) *ImageLabel {
	return NewImageLabelWithImageSize(img, Size{})
}

// NewImageLabelWithImageSize creates a new ImageLabel with the specified image. The image will be
// set to the specified size.
func NewImageLabelWithImageSize(img *Image, size Size) *ImageLabel {
	label := &ImageLabel{image: img}
	if size.Width <= 0 || size.Height <= 0 {
		label.SetSizer(label)
	} else {
		label.SetSizer(&imageLabelSizer{label: label, size: size})
	}
	label.SetPaintHandler(label)
	return label
}

// Sizes implements Sizer
func (label *ImageLabel) Sizes(hint Size) (min, pref, max Size) {
	size := label.image.Size()
	if border := label.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, size
}

// OnPaint implements PaintHandler
func (label *ImageLabel) OnPaint(g Graphics, dirty Rect) {
	bounds := label.LocalInsetBounds()
	size := label.image.Size()
	if size.Width < bounds.Width {
		bounds.X += (bounds.Width - size.Width) / 2
		bounds.Width = size.Width
	}
	if size.Height < bounds.Height {
		bounds.Y += (bounds.Height - size.Height) / 2
		bounds.Height = size.Height
	}
	g.DrawImageInRect(label.image, bounds)
}

type imageLabelSizer struct {
	label *ImageLabel
	size  Size
}

// Sizes implements Sizer
func (sizer *imageLabelSizer) Sizes(hint Size) (min, pref, max Size) {
	pref = sizer.size
	if border := sizer.label.Border(); border != nil {
		pref.AddInsets(border.Insets())
	}
	return pref, pref, pref
}