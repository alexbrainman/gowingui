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
	// TODO: probably, do not need that gloable variable
	moduleHandle syscall.Handle
	cursorArrow  syscall.Handle
)

func init() {
	var err error
	moduleHandle, err = uit.M.GetModuleHandle(nil)
	if err != nil {
		panic(err)
	}
	cursorArrow, err = uit.M.LoadCursor(0, winapi.IDC_ARROW)
	if err != nil {
		panic(err)
	}
}

// TODO: do not know how to get rid of class WndProc altogether
func wndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (rc uintptr) {
	return uit.M.DefWindowProc(hwnd, msg, wparam, lparam)
}

// WindowClass represents GUI window class.
type WindowClass struct {
	Name string
	Atom uint16
}

func RegisterWindowClass(name string) (*WindowClass, error) {
	myicon, err := uit.M.LoadIcon(0, winapi.IDI_APPLICATION)
	if err != nil {
		return nil, err
	}
	wc := winapi.Wndclassex{
		WndProc:    syscall.NewCallback(wndProc),
		Instance:   moduleHandle,
		Icon:       myicon,
		Cursor:     cursorArrow,
		Background: winapi.COLOR_BTNFACE + 1,
		ClassName:  syscall.StringToUTF16Ptr(name),
		IconSm:     myicon,
	}
	wc.Size = uint32(unsafe.Sizeof(wc))
	a, err := uit.M.RegisterClassEx(&wc)
	if err != nil {
		return nil, err
	}
	return &WindowClass{
		Name: name,
		Atom: a,
	}, nil
}

// Window represents GUI window.
type Window struct {
	*WinControl
}

func (c *WindowClass) CreateWindow(text string, r image.Rectangle) (*Window, error) {
	// TODO: handle default window position / size with winapi.CW_USEDEFAULT
	h, err := uit.M.CreateWindowEx(
		winapi.WS_EX_CLIENTEDGE,
		syscall.StringToUTF16Ptr(c.Name),
		syscall.StringToUTF16Ptr(text),
		winapi.WS_OVERLAPPEDWINDOW,
		int32(r.Min.X), int32(r.Min.Y), int32(r.Dx()), int32(r.Dy()),
		0,
		0,
		moduleHandle,
		0)
	if err != nil {
		return nil, err
	}
	return &Window{
		WinControl: &WinControl{Handle: h},
	}, nil
}

func (w *Window) ShowWindow(how int32) (wasViaible bool) {
	return uit.M.ShowWindow(w.Handle, how)
}

func (w *Window) Update() error {
	return uit.M.UpdateWindow(w.Handle)
}

func (w *Window) Close() error {
	return uit.M.PostMessage(w.Handle, winapi.WM_CLOSE, 0, 0)
}
