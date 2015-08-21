// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"errors"
	"image"
	"image/color"
	"syscall"

	"github.com/alexbrainman/gowingui/subclass"
	"github.com/alexbrainman/gowingui/uit"
	"github.com/alexbrainman/gowingui/winapi"
)

type PaintCanvas struct {
	HDC             winapi.HDC
	Rect            *winapi.Rect
	EraseBackground bool

	// keep tabs on what needs to be restored
	// at the end of paint transaction.
	restoreTextColor     bool
	textColorToRestoreTo winapi.COLORREF
	restoreBkColor       bool
	bkColorToRestoreTo   winapi.COLORREF
	restoreFont          bool
	fontToRestoreTo      syscall.Handle
}

func (c *WinControl) AddPaint(paint func(pc *PaintCanvas)) error {
	proc := func(p *subclass.Params) uintptr {
		if p.Msg == winapi.WM_PAINT {
			var ps winapi.PaintStruct
			hdc, err := uit.M.BeginPaint(c.Handle, &ps)
			if err != nil {
				panic(err)
			}
			defer uit.M.EndPaint(c.Handle, &ps)

			pc := &PaintCanvas{
				HDC:             hdc,
				Rect:            &ps.Paint,
				EraseBackground: ps.Erase != 0,
			}

			defer func() {
				if pc.restoreTextColor {
					uit.M.SetTextColor(pc.HDC, pc.textColorToRestoreTo)
				}
				if pc.restoreBkColor {
					uit.M.SetBkColor(pc.HDC, pc.bkColorToRestoreTo)
				}
				if pc.restoreFont {
					uit.M.SelectObject(pc.HDC, pc.fontToRestoreTo)
				}
			}()

			paint(pc)
		}
		return p.CallDefaultProc()
	}
	return c.AddHandler(proc)
}

// TextOut output string text at position p onto canvas pc.
// It uses currently selected font and font color.
func (pc *PaintCanvas) TextOut(p image.Point, text string) error {
	buf := syscall.StringToUTF16(text)
	buf = buf[:len(buf)-1] // remove terminating 0
	return uit.M.TextOut(pc.HDC, int32(p.X), int32(p.Y), &buf[0], int32(len(buf)))
}

// SetTextColor selects color c to be used for any future text output to canvas pc.
// It returns previous text color (color before the change).
// It also records "original" canvas text color, so it can be used to restore
// canvas text color at the end of paint transaction (se per windows api).
func (pc *PaintCanvas) SetTextColor(c color.Color) (prevc Color, err error) {
	cr := winapi.COLORREF(NewColor(c))
	pcr := uit.M.SetTextColor(pc.HDC, cr)
	if pcr == winapi.CLR_INVALID {
		return Color(pcr), errors.New("SetTextColor: Invalid color")
	}
	if !pc.restoreTextColor {
		pc.restoreTextColor = true
		pc.textColorToRestoreTo = pcr
	}
	return Color(pcr), nil
}

func (pc *PaintCanvas) SetBkColor(c color.Color) (prevc Color, err error) {
	cr := winapi.COLORREF(NewColor(c))
	pcr := uit.M.SetBkColor(pc.HDC, cr)
	if pcr == winapi.CLR_INVALID {
		return Color(pcr), errors.New("SetBkColor: Invalid color")
	}
	if !pc.restoreBkColor {
		pc.restoreBkColor = true
		pc.bkColorToRestoreTo = pcr
	}
	return Color(pcr), nil
}

func (pc *PaintCanvas) SetFont(f *Font) error {
	pf, err := uit.M.SelectObject(pc.HDC, f.Handle)
	if err != nil {
		return err
	}
	if !pc.restoreFont {
		pc.restoreFont = true
		pc.fontToRestoreTo = pf
	}
	return nil
}
