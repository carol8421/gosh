// $G $D/$F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main	
 	
func main() {
	x := 0;
	x = ^x;  // unary ^ not yet implemented
	if x != ^0 { panic(x, " ", ^0) }
}

/*
uetli:~/Source/go/test/bugs gri$ 6g bug082.go 
bug082.go:7: fatal error: optoas: no entry COM-<int32>INT32
*/
