package internal

// Varint aliases PutVarint.
func Varint(buf []byte) (int, int) {
	var out int
	n := PutVarint(buf, &out)
	return out, n
}

// PutVarint decodes the received varint encoded byte slice, returning
// the number of bytes used. The decoded integer is placed in out.
//
// The description of the SQLite3 varint spec is copied from the docs
// for ease of understanding:
//
//   A variable-length integer or "varint" is a static Huffman encoding of
//   64-bit twos-complement integers that uses less space for small positive
//   values. A varint is between 1 and 9 bytes in length. The varint
//   consists of either zero or more bytes which have the high-order bit set
//   followed by a single byte with the high-order bit clear, or nine bytes,
//   whichever is shorter. The lower seven bits of each of the first eight
//   bytes and all 8 bits of the ninth byte are used to reconstruct the
//   64-bit twos-complement integer. Varints are big-endian: bits taken from
//   the earlier byte of the varint are more significant than bits taken
//   from the later bytes.
//
//   Source: https://www.sqlite.org/fileformat2.html
//
// Note: SQLite4 proposed a variation of the varint used in SQLite3:
// http://www.sqlite.org/src4/doc/trunk/www/varint.wiki
// So keep in mind that SQLite3 uses a different implementation, as
// implemented below.
func PutVarint(buf []byte, out *int) int {
	// (bX & 0x80) is non-zero if the high-bit of x is set. For SQLite3's varint,
	// we select all bytes from buf[0] to buf[y] where (buf[y] & 0x80 == 0).

	// 1-byte varint: [0, 128)
	b0 := buf[0]
	if b0&0x80 == 0 {
		*out = int(b0)
		return 1
	}

	// 2-byte varint: [128, 16384)
	b1 := buf[1]
	if b1&0x80 == 0 {
		// We only care about the lower 7 bits of each byte, so remove the high-bit and shift.
		*out = int(b0&0x7f)<<7 | int(b1)
		return 2
	}

	// 3-byte varint: [16384, 2097152)
	b2 := buf[2]
	if b2&0x80 == 0 {
		*out = int(b0&0x7f)<<14 | int(b1&0x7f)<<7 | int(b2)
		return 3
	}

	// 4-byte varint: [2097152, 268435456)
	b3 := buf[3]
	if b3&0x80 == 0 {
		*out = int(b0&0x7f)<<21 | int(b1&0x7f)<<14 | int(b2&0x7f)<<7 | int(b3)
		return 4
	}

	// 5-byte varint: [268435456, 34359738368)
	b4 := buf[4]
	if b4&0x80 == 0 {
		*out = int(b0&0x7f)<<28 | int(b1&0x7f)<<21 | int(b2&0x7f)<<14 | int(b3&0x7f)<<7 | int(b4)
		return 5
	}

	// 6-byte varint: [34359738368, 4398046511104)
	b5 := buf[5]
	if b5&0x80 == 0 {
		*out = int(b0&0x7f)<<35 | int(b1&0x7f)<<28 | int(b2&0x7f)<<21 | int(b3&0x7f)<<14 | int(b4&0x7f)<<7 | int(b5)
		return 6
	}

	// 7-byte varint: [4398046511104, 562949953421312)
	b6 := buf[6]
	if b6&0x80 == 0 {
		*out = int(b0&0x7f)<<42 | int(b1&0x7f)<<35 | int(b2&0x7f)<<28 | int(b3&0x7f)<<21 | int(b4&0x7f)<<14 | int(b5&0x7f)<<7 | int(b6)
		return 7
	}

	// 8-byte varint: [562949953421312, 7.205759404e16)
	b7 := buf[7]
	if b7&0x80 == 0 {
		*out = int(b0&0x7f)<<49 | int(b1&0x7f)<<42 | int(b2&0x7f)<<35 | int(b3&0x7f)<<28 | int(b4&0x7f)<<21 | int(b5&0x7f)<<14 | int(b6&0x7f)<<7 | int(b7)
		return 8
	}

	b8 := buf[8]
	*out = int(b0&0x7f)<<56 | int(b1&0x7f)<<49 | int(b2&0x7f)<<42 | int(b3&0x7f)<<35 | int(b4&0x7f)<<28 | int(b5&0x7f)<<21 | int(b6&0x7f)<<14 | int(b7&0x7f)<<7 | int(b8)
	return 9
}
