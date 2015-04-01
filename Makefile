# Copyright 2011 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GOOS=windows

all: gowingui.exe

gowingui.exe:
	go build -o $@ -ldflags -Hwindowsgui

install:
	go install -ldflags -Hwindowsgui

clean:
	rm -f gowingui.exe

zwinapi.go: winapi.go
	(echo '// +build windows'; go run $(GOROOT)/src/syscall/mksyscall_windows.go -- $<) > $@
