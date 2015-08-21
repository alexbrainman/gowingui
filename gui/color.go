// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"image/color"

	"github.com/alexbrainman/gowingui/winapi"
)

// Color represents gui color.
type Color winapi.COLORREF

// RGB creates Color out of its R, G,and B components.
func RGB(r, g, b byte) Color {
	return Color(winapi.RGB(r, g, b))
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(winapi.GetRValue(winapi.COLORREF(c)))
	g = uint32(winapi.GetGValue(winapi.COLORREF(c)))
	b = uint32(winapi.GetBValue(winapi.COLORREF(c)))
	// no alpha value in COLORREF
	return
}

// NewColor creates Color out of color.Color c.
func NewColor(c color.Color) Color {
	r, g, b, _ := c.RGBA()
	return RGB(byte(r/0x100), byte(g/0x100), byte(b/0x100))
}
