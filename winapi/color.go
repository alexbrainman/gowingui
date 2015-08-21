// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

const (
	CLR_NONE    = 0xffffffff
	CLR_INVALID = CLR_NONE
)

type COLORREF uint32

func RGB(r, g, b byte) COLORREF {
	return COLORREF(r) | COLORREF(g)<<8 | COLORREF(b)<<16
}

func GetRValue(c COLORREF) byte {
	return byte(c)
}

func GetGValue(c COLORREF) byte {
	return byte(c >> 8)
}

func GetBValue(c COLORREF) byte {
	return byte(c >> 16)
}
