// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"strconv"

	"github.com/alexbrainman/gowingui/winapi"
)

var msgNames = map[Message]string{
	winapi.WM_CREATE:            "WM_CREATE",
	winapi.WM_DESTROY:           "WM_DESTROY",
	winapi.WM_CLOSE:             "WM_CLOSE",
	winapi.WM_COMMAND:           "WM_COMMAND",
	winapi.WM_KILLFOCUS:         "WM_KILLFOCUS",
	winapi.WM_SETCURSOR:         "WM_SETCURSOR",
	winapi.WM_NCHITTEST:         "WM_NCHITTEST",
	winapi.WM_NCPAINT:           "WM_NCPAINT",
	winapi.WM_MOUSEMOVE:         "WM_MOUSEMOVE",
	winapi.WM_SETTEXT:           "WM_SETTEXT",
	winapi.WM_GETTEXT:           "WM_GETTEXT",
	winapi.WM_GETTEXTLENGTH:     "WM_GETTEXTLENGTH",
	winapi.WM_CTLCOLORBTN:       "WM_CTLCOLORBTN",
	winapi.WM_WINDOWPOSCHANGING: "WM_WINDOWPOSCHANGING",
	winapi.WM_WINDOWPOSCHANGED:  "WM_WINDOWPOSCHANGED",
	winapi.WM_NCACTIVATE:        "WM_NCACTIVATE",
	winapi.WM_ACTIVATE:          "WM_ACTIVATE",
	winapi.WM_ACTIVATEAPP:       "WM_ACTIVATEAPP",
	winapi.WM_NCDESTROY:         "WM_NCDESTROY",
	winapi.WM_MOUSEACTIVATE:     "WM_MOUSEACTIVATE",
	winapi.WM_PARENTNOTIFY:      "WM_PARENTNOTIFY",
	winapi.WM_GETICON:           "WM_GETICON",
	winapi.WM_GETMINMAXINFO:     "WM_GETMINMAXINFO",
	winapi.WM_NCCREATE:          "WM_NCCREATE",
	winapi.WM_NCCALCSIZE:        "WM_NCCALCSIZE",
	winapi.WM_SHOWWINDOW:        "WM_SHOWWINDOW",
	winapi.WM_SETFOCUS:          "WM_SETFOCUS",
	winapi.WM_ERASEBKGND:        "WM_ERASEBKGND",
	winapi.WM_SIZE:              "WM_SIZE",
	winapi.WM_MOVE:              "WM_MOVE",
	winapi.WM_PAINT:             "WM_PAINT",
	winapi.WM_NCMOUSEMOVE:       "WM_NCMOUSEMOVE",
	winapi.WM_IME_SETCONTEXT:    "WM_IME_SETCONTEXT",
	winapi.WM_IME_NOTIFY:        "WM_IME_NOTIFY",
	winapi.WM_CTLCOLORSTATIC:    "WM_CTLCOLORSTATIC",
	winapi.WM_CTLCOLOREDIT:      "WM_CTLCOLOREDIT",
	winapi.WM_SYSCOMMAND:        "WM_SYSCOMMAND",
	winapi.BM_SETSTATE:          "WM_SYSCOMMAND",
	winapi.WM_LBUTTONDOWN:       "WM_LBUTTONDOWN",
	winapi.WM_CAPTURECHANGED:    "WM_CAPTURECHANGED",
	winapi.WM_LBUTTONUP:         "WM_LBUTTONUP",
	winapi.WM_SYSKEYDOWN:        "WM_SYSKEYDOWN",
}

type Message uint32

func (m Message) String() string {
	name := "unknown"
	if s, ok := msgNames[m]; ok {
		name = s
	} else if m >= winapi.WM_USER {
		name = "WM_USER+" + strconv.Itoa(int(m-winapi.WM_USER))
	}
	return name + "(" + strconv.Itoa(int(m)) + ")"
}
