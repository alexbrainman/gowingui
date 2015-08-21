// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"syscall"
	"unsafe"
)

type HDC syscall.Handle

type Wndclassex struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   syscall.Handle
	Icon       syscall.Handle
	Cursor     syscall.Handle
	Background syscall.Handle
	MenuName   *uint16
	ClassName  *uint16
	IconSm     syscall.Handle
}

type Point struct {
	X int32
	Y int32
}

type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type Msg struct {
	Hwnd    syscall.Handle
	Message uint32
	Wparam  uintptr
	Lparam  uintptr
	Time    uint32
	Pt      Point
}

type MinMaxInfo struct {
	Reserved     Point
	MaxSize      Point
	MaxPosition  Point
	MinTrackSize Point
	MaxTrackSize Point
}

type PaintStruct struct {
	HDC       HDC
	Erase     int32
	Paint     Rect
	Restore   int32
	IncUpdate int32
	Reserved  [32]byte
}

const (
	// Window styles
	WS_OVERLAPPED   = 0
	WS_POPUP        = 0x80000000
	WS_CHILD        = 0x40000000
	WS_MINIMIZE     = 0x20000000
	WS_VISIBLE      = 0x10000000
	WS_DISABLED     = 0x8000000
	WS_CLIPSIBLINGS = 0x4000000
	WS_CLIPCHILDREN = 0x2000000
	WS_MAXIMIZE     = 0x1000000
	WS_CAPTION      = WS_BORDER | WS_DLGFRAME
	WS_BORDER       = 0x800000
	WS_DLGFRAME     = 0x400000
	WS_VSCROLL      = 0x200000
	WS_HSCROLL      = 0x100000
	WS_SYSMENU      = 0x80000
	WS_THICKFRAME   = 0x40000
	WS_GROUP        = 0x20000
	WS_TABSTOP      = 0x10000
	WS_MINIMIZEBOX  = 0x20000
	WS_MAXIMIZEBOX  = 0x10000
	WS_TILED        = WS_OVERLAPPED
	WS_ICONIC       = WS_MINIMIZE
	WS_SIZEBOX      = WS_THICKFRAME
	// Common Window Styles
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_TILEDWINDOW      = WS_OVERLAPPEDWINDOW
	WS_POPUPWINDOW      = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_CHILDWINDOW      = WS_CHILD

	WS_EX_CLIENTEDGE = 0x200

	// Some windows messages
	WM_CREATE            = 1
	WM_DESTROY           = 2
	WM_MOVE              = 3
	WM_SIZE              = 5
	WM_ACTIVATE          = 6
	WM_SETFOCUS          = 7
	WM_KILLFOCUS         = 8
	WM_SETTEXT           = 12
	WM_GETTEXT           = 13
	WM_GETTEXTLENGTH     = 14
	WM_PAINT             = 15
	WM_CLOSE             = 16
	WM_QUIT              = 18
	WM_ERASEBKGND        = 20
	WM_SHOWWINDOW        = 24
	WM_ACTIVATEAPP       = 28
	WM_SETCURSOR         = 32
	WM_MOUSEACTIVATE     = 33
	WM_GETMINMAXINFO     = 36
	WM_WINDOWPOSCHANGING = 70
	WM_WINDOWPOSCHANGED  = 71
	WM_GETICON           = 127
	WM_NCCREATE          = 129
	WM_NCDESTROY         = 130
	WM_NCCALCSIZE        = 131
	WM_NCHITTEST         = 132
	WM_NCPAINT           = 133
	WM_NCACTIVATE        = 134
	WM_NCMOUSEMOVE       = 160
	BM_SETSTATE          = 243
	WM_SYSKEYDOWN        = 260
	WM_COMMAND           = 273
	WM_SYSCOMMAND        = 274
	WM_CTLCOLOREDIT      = 307
	WM_CTLCOLORBTN       = 309
	WM_CTLCOLORSTATIC    = 312
	WM_MOUSEMOVE         = 512
	WM_LBUTTONDOWN       = 513
	WM_LBUTTONUP         = 514
	WM_PARENTNOTIFY      = 528
	WM_CAPTURECHANGED    = 533
	WM_IME_SETCONTEXT    = 641
	WM_IME_NOTIFY        = 642
	WM_USER              = 1024

	// Some button control styles
	BS_DEFPUSHBUTTON = 1

	// Some color constants
	COLOR_WINDOW  = 5
	COLOR_BTNFACE = 15

	// Default window position
	// TODO: fix num conversion
	CW_USEDEFAULT = 0x80000000 - 0x100000000

	// Show window default style
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11

	HWND_DESKTOP = syscall.Handle(0)
	HWND_MESSAGE = ^syscall.Handle(0) - 3 + 1

	// GetAncestor flags
	GA_PARENT    = 1
	GA_ROOT      = 2
	GA_ROOTOWNER = 3
)

var (
	// Some globally known cursors
	IDC_ARROW = MakeIntResource(32512)
	IDC_IBEAM = MakeIntResource(32513)
	IDC_WAIT  = MakeIntResource(32514)
	IDC_CROSS = MakeIntResource(32515)

	// Some globally known icons
	IDI_APPLICATION = MakeIntResource(32512)
	IDI_HAND        = MakeIntResource(32513)
	IDI_QUESTION    = MakeIntResource(32514)
	IDI_EXCLAMATION = MakeIntResource(32515)
	IDI_ASTERISK    = MakeIntResource(32516)
	IDI_WINLOGO     = MakeIntResource(32517)
	IDI_WARNING     = IDI_EXCLAMATION
	IDI_ERROR       = IDI_HAND
	IDI_INFORMATION = IDI_ASTERISK
)

// api

//sys	GetModuleHandle(modname *uint16) (handle syscall.Handle, err error) = GetModuleHandleW
//sys	RegisterClassEx(wndclass *Wndclassex) (atom uint16, err error) = user32.RegisterClassExW
//sys	CreateWindowEx(exstyle uint32, classname *uint16, windowname *uint16, style uint32, x int32, y int32, width int32, height int32, wndparent syscall.Handle, menu syscall.Handle, instance syscall.Handle, param uintptr) (hwnd syscall.Handle, err error) = user32.CreateWindowExW
//sys	DefWindowProc(hwnd syscall.Handle, msg uint32, wparam uintptr, lparam uintptr) (lresult uintptr) = user32.DefWindowProcW
//sys	DestroyWindow(hwnd syscall.Handle) (err error) = user32.DestroyWindow
//sys	PostQuitMessage(exitcode int32) = user32.PostQuitMessage
//sys	ShowWindow(hwnd syscall.Handle, cmdshow int32) (wasvisible bool) = user32.ShowWindow
//sys	IsWindowVisible(hwnd syscall.Handle) (visible bool) = user32.IsWindowVisible
//sys	UpdateWindow(hwnd syscall.Handle) (err error) = user32.UpdateWindow
//sys	GetMessage(msg *Msg, hwnd syscall.Handle, MsgFilterMin uint32, MsgFilterMax uint32) (ret int32, err error) [failretval==-1] = user32.GetMessageW
//sys	TranslateMessage(msg *Msg) (done bool) = user32.TranslateMessage
//sys	DispatchMessage(msg *Msg) (ret int32) = user32.DispatchMessageW
//sys	LoadIcon(instance syscall.Handle, iconname *uint16) (icon syscall.Handle, err error) = user32.LoadIconW
//sys	LoadCursor(instance syscall.Handle, cursorname *uint16) (cursor syscall.Handle, err error) = user32.LoadCursorW
//sys	SetCursor(cursor syscall.Handle) (precursor syscall.Handle, err error) = user32.SetCursor
//sys	SendMessage(hwnd syscall.Handle, msg uint32, wparam uintptr, lparam uintptr) (lresult uintptr) = user32.SendMessageW
//sys	PostMessage(hwnd syscall.Handle, msg uint32, wparam uintptr, lparam uintptr) (err error) = user32.PostMessageW
//sys	GetWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) = user32.GetWindowTextW
//sys	GetWindowTextLength(hwnd syscall.Handle) (len int32, err error) = user32.GetWindowTextLengthW
//sys	SetWindowText(hwnd syscall.Handle, str *uint16) (err error) = user32.SetWindowTextW
//sys	EnableWindow(hwnd syscall.Handle, enable bool) (wasenabled bool) = user32.EnableWindow
//sys	IsWindowEnabled(hwnd syscall.Handle) (enabled bool) = user32.IsWindowEnabled
//sys	GetCursorPos(point *Point) (err error) = user32.GetCursorPos

//sys	SetWindowSubclass(hwnd syscall.Handle, fn uintptr, id uintptr, refData *uint32) (err error) = comctl32.SetWindowSubclass
//sys	DefSubclassProc(hwnd syscall.Handle, msg uint32, wparam uintptr, lparam uintptr) (lresult uintptr) = comctl32.DefSubclassProc
//sys	RemoveWindowSubclass(hwnd syscall.Handle, fn uintptr, id uintptr) (err error) = comctl32.RemoveWindowSubclass

//sys	GetCurrentThreadId() (id uint32)

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}
