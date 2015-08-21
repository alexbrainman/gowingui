// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"image"
	"syscall"

	"github.com/alexbrainman/gowingui/uit"
	"github.com/alexbrainman/gowingui/winapi"
)

// Edit represents edit standard control.
type Edit struct {
	*WinControl
}

func (c *WinControl) AddEdit(id int, text string, r image.Rectangle) (*Edit, error) {
	h, err := uit.M.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr("EDIT"),
		syscall.StringToUTF16Ptr(text),
		winapi.WS_CHILD|winapi.WS_VISIBLE,
		int32(r.Min.X), int32(r.Min.Y), int32(r.Dx()), int32(r.Dy()),
		c.Handle, syscall.Handle(id), moduleHandle, 0)
	if err != nil {
		return nil, err
	}
	return &Edit{
		WinControl: &WinControl{Handle: h},
	}, nil
}
