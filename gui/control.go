// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gui

import (
	"image"
	"runtime"
	"syscall"

	"github.com/alexbrainman/gowingui/subclass"
	"github.com/alexbrainman/gowingui/uit"
	"github.com/alexbrainman/gowingui/winapi"
)

var desktop = uit.M.GetDesktopWindow()

// Control
type Control interface {
	// geometry
	Rect() image.Rectangle
	ClientRect() image.Rectangle
	// text handling
	Text() string
	SetText(text string)
	// painting
	Invalidate() error
	InvalidateRect(r image.Rectangle) error
	// TODO: add other methods
}

// WinControl
type WinControl struct {
	Handle   syscall.Handle
	handlers []*subclass.Proc
}

// Windows message handlers.

func (c *WinControl) AddHandler(p func(*subclass.Params) uintptr) error {
	proc, err := subclass.New(c.Handle, p)
	if err != nil {
		return err
	}
	c.handlers = append(c.handlers, proc)
	if len(c.handlers) == 1 {
		runtime.SetFinalizer(c, (*WinControl).Release)
	}
	return nil
}

func (c *WinControl) Release() {
	runtime.SetFinalizer(c, nil)
	for _, proc := range c.handlers {
		proc.Remove()
	}
	c.handlers = nil
}

func (c *WinControl) AddClick(click func()) error {
	proc := func(p *subclass.Params) uintptr {
		if p.Msg == winapi.WM_LBUTTONUP {
			click()
		}
		return p.CallDefaultProc()
	}
	return c.AddHandler(proc)
}

func (c *WinControl) AddResize(resize func(w, h int)) error {
	proc := func(p *subclass.Params) uintptr {
		if p.Msg == winapi.WM_SIZE {
			w := uint16(p.Lparam)
			h := uint16(p.Lparam >> 16)
			resize(int(w), int(h))
		}
		return p.CallDefaultProc()
	}
	return c.AddHandler(proc)
}

// Destroy closes window w and release all os resources associated with it.
func (w *Window) Destroy() error {
	return uit.M.DestroyWindow(w.Handle)
}

// Get/Set control text.

func (c *WinControl) Text() string {
	n, err := uit.M.GetWindowTextLength(c.Handle)
	if err != nil {
		//panic(err)
		return "" // ignore error
	}
	buf := make([]uint16, n+1)
	_, err = uit.M.GetWindowText(c.Handle, &buf[0], int32(len(buf)))
	if err != nil {
		//panic(err)
		return "" // ignore error
	}
	return syscall.UTF16ToString(buf)
}

func (c *WinControl) SetText(text string) {
	buf := syscall.StringToUTF16(text)
	uit.M.SetWindowText(c.Handle, &buf[0])
}

// Enable/Disable control.

func (c *WinControl) IsEnabled() bool {
	return uit.M.IsWindowEnabled(c.Handle)
}

func (c *WinControl) Enable() (prev bool) {
	return uit.M.EnableWindow(c.Handle, true)
}

func (c *WinControl) Disable() (prev bool) {
	return uit.M.EnableWindow(c.Handle, false)
}

// Show/Hide control.

func (c *WinControl) IsVisible() bool {
	return uit.M.IsWindowVisible(c.Handle)
}

func (c *WinControl) Show() (prev bool) {
	return uit.M.ShowWindow(c.Handle, winapi.SW_SHOW)
}

func (c *WinControl) Hide() (prev bool) {
	return uit.M.ShowWindow(c.Handle, winapi.SW_HIDE)
}

// Geometry.

// Rect retrieves the dimensions of control's c bounding rectangle.
// The coordinates are measured relative to c.
func (c *WinControl) Rect() image.Rectangle {
	var r winapi.Rect
	err := uit.M.GetWindowRect(c.Handle, &r)
	if err != nil {
		panic(err)
	}
	parent := uit.M.GetAncestor(c.Handle, winapi.GA_PARENT)
	if parent == desktop {
		// desktop is our parent -> just return
		return image.Rect(int(r.Left), int(r.Top), int(r.Right), int(r.Bottom))
	}
	p := []winapi.Point{
		{X: r.Left, Y: r.Top},
		{X: r.Right, Y: r.Bottom},
	}
	err = uit.M.MapWindowPoints(winapi.HWND_DESKTOP, parent, &p[0], uint32(len(p)))
	if err != nil {
		panic(err)
	}
	return image.Rect(int(p[0].X), int(p[0].Y), int(p[1].X), int(p[1].Y))
}

func (c *WinControl) SetRect(r image.Rectangle) error {
	return uit.M.MoveWindow(c.Handle,
		int32(r.Min.X), int32(r.Min.Y), int32(r.Dx()), int32(r.Dy()), true)
}

// ClientRect retrieves the coordinates of control's c client area.
// The coordinates are measured relative to c.
func (c *WinControl) ClientRect() image.Rectangle {
	var r winapi.Rect
	err := uit.M.GetClientRect(c.Handle, &r)
	if err != nil {
		panic(err)
	}
	return image.Rect(int(r.Left), int(r.Top), int(r.Right), int(r.Bottom))
}

// Painting.

func (c *WinControl) Invalidate() error {
	return uit.M.InvalidateRect(c.Handle, nil, true)
}

func (c *WinControl) InvalidateRect(r image.Rectangle) error {
	wr := winapi.Rect{
		Left:   int32(r.Min.X),
		Top:    int32(r.Min.Y),
		Right:  int32(r.Max.X),
		Bottom: int32(r.Max.Y),
	}
	return uit.M.InvalidateRect(c.Handle, &wr, true)
}
