// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gzip package implements reading (and eventually writing) of
// gzip format compressed files, as specified in RFC 1952.
package gzip

import (
	"bufio";
	"compress/flate";
	"hash";
	"hash/crc32";
	"io";
	"os";
)

const (
	gzipID1		= 0x1f;
	gzipID2		= 0x8b;
	gzipDeflate	= 8;
	flagText	= 1<<0;
	flagHdrCrc	= 1<<1;
	flagExtra	= 1<<2;
	flagName	= 1<<3;
	flagComment	= 1<<4;
)

func makeReader(r io.Reader) flate.Reader {
	if rr, ok := r.(flate.Reader); ok {
		return rr;
	}
	return bufio.NewReader(r);
}

var HeaderError os.Error = os.ErrorString("invalid gzip header")
var ChecksumError os.Error = os.ErrorString("gzip checksum error")

// An Inflater is an io.Reader that can be read to retrieve
// uncompressed data from a gzip-format compressed file.
// The gzip file stores a header giving metadata about the compressed file.
// That header is exposed as the fields of the Inflater struct.
//
// In general, a gzip file can be a concatenation of gzip files,
// each with its own header.  Reads from the Inflater
// return the concatenation of the uncompressed data of each.
// Only the first header is recorded in the Inflater fields.
//
// Gzip files store a length and checksum of the uncompressed data.
// The Inflater will return a ChecksumError when Read
// reaches the end of the uncompressed data if it does not
// have the expected length or checksum.  Clients should treat data
// returned by Read as tentative until they receive the successful
// (zero length, nil error) Read marking the end of the data.
type Inflater struct {
	Comment	string;	// comment
	Extra	[]byte;	// "extra data"
	Mtime	uint32;	// modification time (seconds since January 1, 1970)
	Name	string;	// file name
	OS	byte;	// operating system type

	r		flate.Reader;
	inflater	io.ReadCloser;
	digest		hash.Hash32;
	size		uint32;
	flg		byte;
	buf		[512]byte;
	err		os.Error;
	eof		bool;
}

// NewInflater creates a new Inflater reading the given reader.
// The implementation buffers input and may read more data than necessary from r.
// It is the caller's responsibility to call Close on the Inflater when done.
func NewInflater(r io.Reader) (*Inflater, os.Error) {
	z := new(Inflater);
	z.r = makeReader(r);
	z.digest = crc32.NewIEEE();
	if err := z.readHeader(true); err != nil {
		z.err = err;
		return nil, err;
	}
	return z, nil;
}

// GZIP (RFC 1952) is little-endian, unlike ZLIB (RFC 1950).
func get4(p []byte) uint32 {
	return uint32(p[0]) | uint32(p[1])<<8 | uint32(p[2])<<16 | uint32(p[3])<<24;
}

func (z *Inflater) readString() (string, os.Error) {
	var err os.Error;
	for i := 0; ; i++ {
		if i >= len(z.buf) {
			return "", HeaderError;
		}
		z.buf[i], err = z.r.ReadByte();
		if err != nil {
			return "", err;
		}
		if z.buf[i] == 0 {
			return string(z.buf[0:i]), nil;
		}
	}
	panic("not reached");
}

func (z *Inflater) read2() (uint32, os.Error) {
	_, err := z.r.Read(z.buf[0:2]);
	if err != nil {
		return 0, err;
	}
	return uint32(z.buf[0]) | uint32(z.buf[1])<<8, nil;
}

func (z *Inflater) readHeader(save bool) os.Error {
	_, err := io.ReadFull(z.r, z.buf[0:10]);
	if err != nil {
		return err;
	}
	if z.buf[0] != gzipID1 || z.buf[1] != gzipID2 || z.buf[2] != gzipDeflate {
		return HeaderError;
	}
	z.flg = z.buf[3];
	if save {
		z.Mtime = get4(z.buf[4:8]);
		// z.buf[8] is xfl, ignored
		z.OS = z.buf[9];
	}
	z.digest.Reset();
	z.digest.Write(z.buf[0:10]);

	if z.flg & flagExtra != 0 {
		n, err := z.read2();
		if err != nil {
			return err;
		}
		data := make([]byte, n);
		if _, err = io.ReadFull(z.r, data); err != nil {
			return err;
		}
		if save {
			z.Extra = data;
		}
	}

	var s string;
	if z.flg & flagName != 0 {
		if s, err = z.readString(); err != nil {
			return err;
		}
		if save {
			z.Name = s;
		}
	}

	if z.flg & flagComment != 0 {
		if s, err = z.readString(); err != nil {
			return err;
		}
		if save {
			z.Comment = s;
		}
	}

	if z.flg & flagHdrCrc != 0 {
		n, err := z.read2();
		if err != nil {
			return err;
		}
		sum := z.digest.Sum32() & 0xFFFF;
		if n != sum {
			return HeaderError;
		}
	}

	z.digest.Reset();
	z.inflater = flate.NewInflater(z.r);
	return nil;
}

func (z *Inflater) Read(p []byte) (n int, err os.Error) {
	if z.err != nil {
		return 0, z.err;
	}
	if z.eof || len(p) == 0 {
		return 0, nil;
	}

	n, err = z.inflater.Read(p);
	z.digest.Write(p[0:n]);
	z.size += uint32(n);
	if n != 0 || err != os.EOF {
		z.err = err;
		return;
	}

	// Finished file; check checksum + size.
	if _, err := io.ReadFull(z.r, z.buf[0:8]); err != nil {
		z.err = err;
		return 0, err;
	}
	crc32, isize := get4(z.buf[0:4]), get4(z.buf[4:8]);
	sum := z.digest.Sum32();
	if sum != crc32 || isize != z.size {
		z.err = ChecksumError;
		return 0, z.err;
	}

	// File is ok; is there another?
	if err = z.readHeader(false); err != nil {
		z.err = err;
		return;
	}

	// Yes.  Reset and read from it.
	z.digest.Reset();
	z.size = 0;
	return z.Read(p);
}

// Calling Close does not close the wrapped io.Reader originally passed to NewInflater.
func (z *Inflater) Close() os.Error {
	return z.inflater.Close();
}
