// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"syscall"

	"github.com/alexbrainman/gowingui/uit"
	"github.com/alexbrainman/gowingui/winapi"
)

type Font struct {
	Handle syscall.Handle
}

type FontType int32

const (
	Italic FontType = 1 << iota
	Underline
	StrikeOut
)

// OpenFont make font ready to be used. Font face is specified by name parameter.
// Font height can be negative or 0 (see windows documentation). Font weight is
// in the range 0 through 1000, but windows predefined consts can be used (FW_BOLD).
// Font type t is a set of boolean flags for italic, underline and such.
func OpenFont(name string, height int, weight int32, t FontType /* TODO: more parameters */) (f *Font, err error) {
	l := winapi.LOGFONT{
		Height: int32(height),
		Weight: weight,
	}
	if t&Italic != 0 {
		l.Italic = 1
	}
	if t&Underline != 0 {
		l.Underline = 1
	}
	if t&StrikeOut != 0 {
		l.StrikeOut = 1
	}
	copy(l.FaceName[:], syscall.StringToUTF16(name))
	h, err := uit.M.CreateFontIndirect(&l)
	if err != nil {
		return nil, err
	}
	return &Font{Handle: h}, nil
}

// Close release font f.
func (f *Font) Close() error {
	return uit.M.DeleteObject(f.Handle)
}
