// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zlib

import (
	"bytes";
	"io";
	"os";
	"testing";
)

type zlibTest struct {
	desc		string;
	raw		string;
	compressed	[]byte;
	err		os.Error;
}

// Compare-to-golden test data was generated by the ZLIB example program at
// http://www.zlib.net/zpipe.c

var zlibTests = []zlibTest{
	zlibTest{
		"empty",
		"",
		[]byte{0x78, 0x9c, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01},
		nil,
	},
	zlibTest{
		"goodbye",
		"goodbye, world",
		[]byte{
			0x78, 0x9c, 0x4b, 0xcf, 0xcf, 0x4f, 0x49, 0xaa,
			0x4c, 0xd5, 0x51, 0x28, 0xcf, 0x2f, 0xca, 0x49,
			0x01, 0x00, 0x28, 0xa5, 0x05, 0x5e,
		},
		nil,
	},
	zlibTest{
		"bad header",
		"",
		[]byte{0x78, 0x9f, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01},
		HeaderError,
	},
	zlibTest{
		"bad checksum",
		"",
		[]byte{0x78, 0x9c, 0x03, 0x00, 0x00, 0x00, 0x00, 0xff},
		ChecksumError,
	},
	zlibTest{
		"not enough data",
		"",
		[]byte{0x78, 0x9c, 0x03, 0x00, 0x00, 0x00},
		io.ErrUnexpectedEOF,
	},
	zlibTest{
		"excess data is silently ignored",
		"",
		[]byte{
			0x78, 0x9c, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x78, 0x9c, 0xff,
		},
		nil,
	},
}

func TestInflater(t *testing.T) {
	b := new(bytes.Buffer);
	for _, tt := range zlibTests {
		in := bytes.NewBuffer(tt.compressed);
		zlib, err := NewInflater(in);
		if err != nil {
			if err != tt.err {
				t.Errorf("%s: NewInflater: %s", tt.desc, err);
			}
			continue;
		}
		defer zlib.Close();
		b.Reset();
		n, err := io.Copy(b, zlib);
		if err != nil {
			if err != tt.err {
				t.Errorf("%s: io.Copy: %v want %v", tt.desc, err, tt.err);
			}
			continue;
		}
		s := b.String();
		if s != tt.raw {
			t.Errorf("%s: got %d-byte %q want %d-byte %q", tt.desc, n, s, len(tt.raw), tt.raw);
		}
	}
}