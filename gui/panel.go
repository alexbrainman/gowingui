// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"image"
	"syscall"
	"unsafe"

	"github.com/alexbrainman/gowingui/uit"
	"github.com/alexbrainman/gowingui/winapi"
)

var (
	panelClassName = "panel"
)

func init() {
	wc := winapi.Wndclassex{
		WndProc:   syscall.NewCallback(wndProc),
		Instance:  moduleHandle,
		Cursor:    cursorArrow,
		ClassName: syscall.StringToUTF16Ptr(panelClassName),
	}
	wc.Size = uint32(unsafe.Sizeof(wc))
	_, err := uit.M.RegisterClassEx(&wc)
	if err != nil {
		panic(err)
	}
}

// Panel.
type Panel struct {
	*WinControl
}

func (c *WinControl) AddPanel(r image.Rectangle) (*Panel, error) {
	h, err := uit.M.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr(panelClassName),
		nil,
		winapi.WS_CHILD|winapi.WS_VISIBLE|winapi.WS_BORDER,
		int32(r.Min.X), int32(r.Min.Y), int32(r.Dx()), int32(r.Dy()),
		c.Handle, 0, moduleHandle, 0)
	if err != nil {
		return nil, err
	}
	return &Panel{
		WinControl: &WinControl{Handle: h},
	}, nil
}
