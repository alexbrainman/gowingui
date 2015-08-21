# Copyright 2012 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GOOS=windows

TMP=/tmp/mksyscall_windows

zwinapi.go: winapi.go font.go
	go build -o $(TMP) $(GOROOT)/src/syscall/mksyscall_windows.go
	GOOS=windows $(TMP) $^ | gofmt > $@
	rm $(TMP)

