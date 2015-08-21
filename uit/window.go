// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uit

import (
	"syscall"
	"unsafe"

	"github.com/alexbrainman/gowingui/winapi"
)

type call struct {
	trap uintptr
	args []uintptr
	r1   uintptr
	r2   uintptr
	err  syscall.Errno
}

func (c *call) execute() {
	switch l := uintptr(len(c.args)); l {
	case 0:
		c.r1, c.r2, c.err = syscall.Syscall(c.trap, l, 0, 0, 0)
	case 1:
		c.r1, c.r2, c.err = syscall.Syscall(c.trap, l, c.args[0], 0, 0)
	case 2:
		c.r1, c.r2, c.err = syscall.Syscall(c.trap, l, c.args[0], c.args[1], 0)
	case 3:
		c.r1, c.r2, c.err = syscall.Syscall(c.trap, l, c.args[0], c.args[1], c.args[2])
	case 4:
		c.r1, c.r2, c.err = syscall.Syscall6(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], 0, 0)
	case 5:
		c.r1, c.r2, c.err = syscall.Syscall6(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], 0)
	case 6:
		c.r1, c.r2, c.err = syscall.Syscall6(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5])
	case 7:
		c.r1, c.r2, c.err = syscall.Syscall9(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], 0, 0)
	case 8:
		c.r1, c.r2, c.err = syscall.Syscall9(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], c.args[7], 0)
	case 9:
		c.r1, c.r2, c.err = syscall.Syscall9(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], c.args[7], c.args[8])
	case 10:
		c.r1, c.r2, c.err = syscall.Syscall12(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], c.args[7], c.args[8], c.args[9], 0, 0)
	case 11:
		c.r1, c.r2, c.err = syscall.Syscall12(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], c.args[7], c.args[8], c.args[9], c.args[10], 0)
	case 12:
		c.r1, c.r2, c.err = syscall.Syscall12(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], c.args[7], c.args[8], c.args[9], c.args[10], c.args[11])
	case 13:
		c.r1, c.r2, c.err = syscall.Syscall15(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], c.args[7], c.args[8], c.args[9], c.args[10], c.args[11], c.args[12], 0, 0)
	case 14:
		c.r1, c.r2, c.err = syscall.Syscall15(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], c.args[7], c.args[8], c.args[9], c.args[10], c.args[11], c.args[12], c.args[13], 0)
	case 15:
		c.r1, c.r2, c.err = syscall.Syscall15(c.trap, l, c.args[0], c.args[1], c.args[2], c.args[3], c.args[4], c.args[5], c.args[6], c.args[7], c.args[8], c.args[9], c.args[10], c.args[11], c.args[12], c.args[13], c.args[14])
	default:
		panic("(*call).execute: too many aruments")
	}
}

// TODO: perhaps this can be made more efficient
func (t *Thread) syscall(trap uintptr, args []uintptr) (r1, r2 uintptr, err syscall.Errno) {
	c := &call{
		trap: trap,
		args: args,
	}
	t.execute(c)
	return c.r1, c.r2, c.err
}

func (t *Thread) execute(c *call) {
	// TODO: maybe replace GetCurrentThreadId call with something else - it is exensive to call it for every syscall processed.
	if t.tid == winapi.GetCurrentThreadId() {
		c.execute()
	} else {
		winapi.SendMessage(t.iw, winapi.WM_USER, uintptr(unsafe.Pointer(c)), 0)
	}
}

func wndproc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (rc uintptr) {
	switch msg {
	case winapi.WM_USER:
		(*call)(unsafe.Pointer(wparam)).execute()
	case winapi.WM_DESTROY:
		winapi.PostQuitMessage(0)
	default:
		rc = winapi.DefWindowProc(hwnd, msg, wparam, lparam)
	}
	return
}

func makeInvisibleWindow() (syscall.Handle, error) {
	mh, _ := winapi.GetModuleHandle(nil)
	// TODO: make a good class name
	cname := syscall.StringToUTF16Ptr("MAKE-SOME-GOOD-NAME-HERE")
	wc := winapi.Wndclassex{
		WndProc:   syscall.NewCallback(wndproc),
		Instance:  mh,
		ClassName: cname,
	}
	wc.Size = uint32(unsafe.Sizeof(wc))
	_, err := winapi.RegisterClassEx(&wc)
	if err != nil {
		return 0, err
	}
	h, err := winapi.CreateWindowEx(
		0,
		cname,
		nil,
		0,
		0, 0, 0, 0,
		winapi.HWND_MESSAGE,
		0,
		mh,
		0)
	if err != nil {
		return 0, err
	}
	return h, nil
}
