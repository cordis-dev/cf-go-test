// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests for the atomic checker.

package abomaA

func SetX(x int) {
	x = x // self-assignment of x to x
}
