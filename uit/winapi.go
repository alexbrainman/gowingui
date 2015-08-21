// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uit

//sys	GetModuleHandle(modname *uint16) (handle syscall.Handle, err error) = GetModuleHandleW
//sys	CreateWindowEx(exstyle uint32, classname *uint16, windowname *uint16, style uint32, x int32, y int32, width int32, height int32, wndparent syscall.Handle, menu syscall.Handle, instance syscall.Handle, param uintptr) (hwnd syscall.Handle, err error) = user32.CreateWindowExW
//sys	GetCursorPos(point *winapi.Point) (err error) = user32.GetCursorPos
//sys	GetClientRect(hwnd syscall.Handle, rect *winapi.Rect) (err error) = user32.GetClientRect
//sys	GetWindowRect(hwnd syscall.Handle, rect *winapi.Rect) (err error) = user32.GetWindowRect
//sys	MoveWindow(hwnd syscall.Handle, x int32, y int32, w int32, h int32, repaint bool) (err error) = user32.MoveWindow
//sys	InvalidateRect(hwnd syscall.Handle, rect *winapi.Rect, erase bool) (err error) = user32.InvalidateRect
//sys	MapWindowPoints(from syscall.Handle, to syscall.Handle, points *winapi.Point, count uint32) (err error) = user32.MapWindowPoints
//sys	GetWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) = user32.GetWindowTextW
//sys	GetWindowTextLength(hwnd syscall.Handle) (len int32, err error) = user32.GetWindowTextLengthW
//sys	SetWindowText(hwnd syscall.Handle, str *uint16) (err error) = user32.SetWindowTextW
//sys	EnableWindow(hwnd syscall.Handle, enable bool) (wasenabled bool) = user32.EnableWindow
//sys	IsWindowEnabled(hwnd syscall.Handle) (enabled bool) = user32.IsWindowEnabled
//sys	ShowWindow(hwnd syscall.Handle, cmdshow int32) (wasvisible bool) = user32.ShowWindow
//sys	IsWindowVisible(hwnd syscall.Handle) (visible bool) = user32.IsWindowVisible
//sys	LoadIcon(instance syscall.Handle, iconname *uint16) (icon syscall.Handle, err error) = user32.LoadIconW
//sys	LoadCursor(instance syscall.Handle, cursorname *uint16) (cursor syscall.Handle, err error) = user32.LoadCursorW
//sys	RegisterClassEx(wndclass *winapi.Wndclassex) (atom uint16, err error) = user32.RegisterClassExW
//sys	DestroyWindow(hwnd syscall.Handle) (err error) = user32.DestroyWindow
//sys	PostQuitMessage(exitcode int32) = user32.PostQuitMessage
//sys	DefWindowProc(hwnd syscall.Handle, msg uint32, wparam uintptr, lparam uintptr) (lresult uintptr) = user32.DefWindowProcW
//sys	UpdateWindow(hwnd syscall.Handle) (err error) = user32.UpdateWindow
//sys	PostMessage(hwnd syscall.Handle, msg uint32, wparam uintptr, lparam uintptr) (err error) = user32.PostMessageW
//sys	GetAncestor(hwnd syscall.Handle, flags uint32) (ancestor syscall.Handle) = user32.GetAncestor
//sys	GetDesktopWindow() (desktop syscall.Handle) = user32.GetDesktopWindow

//sys	SetWindowSubclass(hwnd syscall.Handle, fn uintptr, id uintptr, refData *uint32) (err error) = comctl32.SetWindowSubclass
//sys	DefSubclassProc(hwnd syscall.Handle, msg uint32, wparam uintptr, lparam uintptr) (lresult uintptr) = comctl32.DefSubclassProc
//sys	RemoveWindowSubclass(hwnd syscall.Handle, fn uintptr, id uintptr) (err error) = comctl32.RemoveWindowSubclass

//sys	BeginPaint(hwnd syscall.Handle, paint *winapi.PaintStruct) (hdc winapi.HDC, err error) = user32.BeginPaint
//sys	EndPaint(hwnd syscall.Handle, paint *winapi.PaintStruct) (err error) = user32.EndPaint
//sys	TextOut(hdc winapi.HDC, x int32, y int32, str *uint16, strlen int32) (err error) = gdi32.TextOutW
//sys	SetTextColor(hdc winapi.HDC, color winapi.COLORREF) (prevcolor winapi.COLORREF) = gdi32.SetTextColor
//sys	SetBkColor(hdc winapi.HDC, color winapi.COLORREF) (prevcolor winapi.COLORREF) = gdi32.SetBkColor
//sys	SelectObject(hdc winapi.HDC, obj syscall.Handle) (handle syscall.Handle, err error) = gdi32.SelectObject
//sys	CreateFontIndirect(font *winapi.LOGFONT) (fh syscall.Handle, err error) = gdi32.CreateFontIndirectW
//sys	DeleteObject(obj syscall.Handle) (err error) = gdi32.DeleteObject
