// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package implements the SHA1 hash algorithm as defined in RFC 3174.
package sha1

import (
	"hash";
	"os";
)

// The size of a SHA1 checksum in bytes.
const Size = 20

const (
	_Chunk	= 64;
	_Init0	= 0x67452301;
	_Init1	= 0xEFCDAB89;
	_Init2	= 0x98BADCFE;
	_Init3	= 0x10325476;
	_Init4	= 0xC3D2E1F0;
)

// digest represents the partial evaluation of a checksum.
type digest struct {
	h	[5]uint32;
	x	[_Chunk]byte;
	nx	int;
	len	uint64;
}

func (d *digest) Reset() {
	d.h[0] = _Init0;
	d.h[1] = _Init1;
	d.h[2] = _Init2;
	d.h[3] = _Init3;
	d.h[4] = _Init4;
	d.nx = 0;
	d.len = 0;
}

// New returns a Hash computing the SHA1 checksum.
func New() hash.Hash {
	d := new(digest);
	d.Reset();
	return d;
}

func (d *digest) Size() int {
	return Size;
}

func (d *digest) Write(p []byte) (nn int, err os.Error) {
	nn = len(p);
	d.len += uint64(nn);
	if d.nx > 0 {
		n := len(p);
		if n > _Chunk - d.nx {
			n = _Chunk - d.nx;
		}
		for i := 0; i < n; i++ {
			d.x[d.nx + i] = p[i];
		}
		d.nx += n;
		if d.nx == _Chunk {
			_Block(d, &d.x);
			d.nx = 0;
		}
		p = p[n:len(p)];
	}
	n := _Block(d, p);
	p = p[n:len(p)];
	if len(p) > 0 {
		for i := 0; i < len(p); i++ {
			d.x[i] = p[i];
		}
		d.nx = len(p);
	}
	return;
}

func (d *digest) Sum() []byte {
	// Padding.  Add a 1 bit and 0 bits until 56 bytes mod 64.
	len := d.len;
	var tmp [64]byte;
	tmp[0] = 0x80;
	if len%64 < 56 {
		d.Write(tmp[0 : 56 - len%64]);
	} else {
		d.Write(tmp[0 : 64 + 56 - len%64]);
	}

	// Length in bits.
	len <<= 3;
	for i := uint(0); i < 8; i++ {
		tmp[i] = byte(len>>(56 - 8*i));
	}
	d.Write(tmp[0:8]);

	if d.nx != 0 {
		panicln("oops");
	}

	p := make([]byte, 20);
	j := 0;
	for i := 0; i < 5; i++ {
		s := d.h[i];
		p[j] = byte(s>>24);
		j++;
		p[j] = byte(s>>16);
		j++;
		p[j] = byte(s>>8);
		j++;
		p[j] = byte(s);
		j++;
	}
	return p;
}
