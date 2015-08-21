// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subclass

import (
	"sync"
	"syscall"
	"unsafe"

	"github.com/alexbrainman/gowingui/uit"
)

type Params struct {
	Hwnd   syscall.Handle
	Msg    uint32
	Wparam uintptr
	Lparam uintptr
	Id     uintptr
	Ref    *uint32
}

func (p *Params) CallDefaultProc() (rc uintptr) {
	return uit.M.DefSubclassProc(p.Hwnd, p.Msg, p.Wparam, p.Lparam)
}

type Proc struct {
	h    syscall.Handle
	id   uintptr
	proc func(*Params) uintptr
}

var (
	mu  sync.Mutex
	gid int
)

func New(h syscall.Handle, proc func(*Params) uintptr) (*Proc, error) {
	mu.Lock()
	p := Proc{h: h, id: uintptr(gid), proc: proc}
	gid++
	mu.Unlock()
	err := uit.M.SetWindowSubclass(h, wndProcUintptr, p.id, (*uint32)(unsafe.Pointer(&p)))
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// TODO: decide how to use Remove everywhere - it is used nowhere at this moment !!!!!!!!!!!!!!!!!!!!!!!!!

func (p *Proc) Remove() error {
	return uit.M.RemoveWindowSubclass(p.h, wndProcUintptr, p.id)
}

var wndProcUintptr = syscall.NewCallback(wndProc)

func wndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr, id uintptr, ref *uint32) (rc uintptr) {
	if ref == nil {
		return uit.M.DefSubclassProc(hwnd, msg, wparam, lparam)
	}
	p := &Params{
		Hwnd:   hwnd,
		Msg:    msg,
		Wparam: wparam,
		Lparam: lparam,
		Id:     id,
		Ref:    ref,
	}
	return (*Proc)(unsafe.Pointer(ref)).proc(p)
}
