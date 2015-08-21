// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uit

import (
	"runtime"
	"syscall"

	"github.com/alexbrainman/gowingui/winapi"
)

type Thread struct {
	tid uint32
	ec  chan *exitStatus
	iw  syscall.Handle
}

type startStatus struct {
	t   *Thread
	err error
}

type exitStatus struct {
	rc  int
	err error
}

func runLoop(c chan<- *startStatus) {
	runtime.LockOSThread()

	iw, err := makeInvisibleWindow()
	if err != nil {
		c <- &startStatus{err: err}
		return
	}
	t := &Thread{
		tid: winapi.GetCurrentThreadId(),
		ec:  make(chan *exitStatus, 1), // so runLoop does not block
		iw:  iw,
	}
	c <- &startStatus{t: t}

	var m winapi.Msg
	for {
		r, err := winapi.GetMessage(&m, 0, 0, 0)
		if err != nil {
			t.ec <- &exitStatus{0, err}
			return
		}
		if r == 0 {
			println(r, err)
		}
		if r == 0 && m.Message == winapi.WM_QUIT {
			break
		}
		winapi.TranslateMessage(&m)
		winapi.DispatchMessage(&m)
	}
	t.ec <- &exitStatus{int(m.Wparam), nil}
}

// TODO: change this whole interface to someting simple (do not use goroutines here) !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func Start() (*Thread, error) {
	c := make(chan *startStatus)
	go runLoop(c)
	ss := <-c
	if ss.err != nil {
		return nil, ss.err
	}
	return ss.t, nil
}

func (t *Thread) Stop() error {
	return winapi.PostMessage(t.iw, winapi.WM_CLOSE, 0, 0)
}

func (t *Thread) Wait() (int, error) {
	es := <-t.ec
	if es.err != nil {
		return 0, es.err
	}
	return es.rc, nil
}
