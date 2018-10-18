// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SHA1 hash algorithm.  See RFC 3174.

package sha1

import (
	"fmt";
	"io";
	"testing";
)

type sha1Test struct {
	out	string;
	in	string;
}

var golden = []sha1Test{
	sha1Test{"da39a3ee5e6b4b0d3255bfef95601890afd80709", ""},
	sha1Test{"86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", "a"},
	sha1Test{"da23614e02469a0d7c7bd1bdab5c9c474b1904dc", "ab"},
	sha1Test{"a9993e364706816aba3e25717850c26c9cd0d89d", "abc"},
	sha1Test{"81fe8bfe87576c3ecb22426f8e57847382917acf", "abcd"},
	sha1Test{"03de6c570bfe24bfc328ccd7ca46b76eadaf4334", "abcde"},
	sha1Test{"1f8ac10f23c5b5bc1167bda84b833e5c057a77d2", "abcdef"},
	sha1Test{"2fb5e13419fc89246865e7a324f476ec624e8740", "abcdefg"},
	sha1Test{"425af12a0743502b322e93a015bcf868e324d56a", "abcdefgh"},
	sha1Test{"c63b19f1e4c8b5f76b25c49b8b87f57d8e4872a1", "abcdefghi"},
	sha1Test{"d68c19a0a345b7eab78d5e11e991c026ec60db63", "abcdefghij"},
	sha1Test{"ebf81ddcbe5bf13aaabdc4d65354fdf2044f38a7", "Discard medicine more than two years old."},
	sha1Test{"e5dea09392dd886ca63531aaa00571dc07554bb6", "He who has a shady past knows that nice guys finish last."},
	sha1Test{"45988f7234467b94e3e9494434c96ee3609d8f8f", "I wouldn't marry him with a ten foot pole."},
	sha1Test{"55dee037eb7460d5a692d1ce11330b260e40c988", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	sha1Test{"b7bc5fb91080c7de6b582ea281f8a396d7c0aee8", "The days of the digital watch are numbered.  -Tom Stoppard"},
	sha1Test{"c3aed9358f7c77f523afe86135f06b95b3999797", "Nepal premier won't resign."},
	sha1Test{"6e29d302bf6e3a5e4305ff318d983197d6906bb9", "For every action there is an equal and opposite government program."},
	sha1Test{"597f6a540010f94c15d71806a99a2c8710e747bd", "His money is twice tainted: 'taint yours and 'taint mine."},
	sha1Test{"6859733b2590a8a091cecf50086febc5ceef1e80", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	sha1Test{"514b2630ec089b8aee18795fc0cf1f4860cdacad", "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	sha1Test{"c5ca0d4a7b6676fc7aa72caa41cc3d5df567ed69", "size:  a.out:  bad magic"},
	sha1Test{"74c51fa9a04eadc8c1bbeaa7fc442f834b90a00a", "The major problem is with sendmail.  -Mark Horton"},
	sha1Test{"0b4c4ce5f52c3ad2821852a8dc00217fa18b8b66", "Give me a rock, paper and scissors and I will move the world.  CCFestoon"},
	sha1Test{"3ae7937dd790315beb0f48330e8642237c61550a", "If the enemy is within range, then so are you."},
	sha1Test{"410a2b296df92b9a47412b13281df8f830a9f44b", "It's well we cannot hear the screams/That we create in others' dreams."},
	sha1Test{"841e7c85ca1adcddbdd0187f1289acb5c642f7f5", "You remind me of a TV show, but that's all right: I watch it anyway."},
	sha1Test{"163173b825d03b952601376b25212df66763e1db", "C is as portable as Stonehedge!!"},
	sha1Test{"32b0377f2687eb88e22106f133c586ab314d5279", "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
	sha1Test{"0885aaf99b569542fd165fa44e322718f4a984e0", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
	sha1Test{"6627d6904d71420b0bf3886ab629623538689f45", "How can you write a big system without C++?  -Paul Glick"},
}

func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i];
		c := New();
		for j := 0; j < 2; j++ {
			io.WriteString(c, g.in);
			s := fmt.Sprintf("%x", c.Sum());
			if s != g.out {
				t.Errorf("sha1[%d](%s) = %s want %s", j, g.in, s, g.out);
				t.FailNow();
			}
			c.Reset();
		}
	}
}