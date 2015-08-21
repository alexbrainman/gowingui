// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uit

var (
	// TODO: perhaps M is not such a good name for "main thread"
	M *Thread
)

func init() {
	var err error
	M, err = Start()
	if err != nil {
		panic("Could not create uit.M: " + err.Error())
	}
}
