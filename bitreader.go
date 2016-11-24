package vorbis

type bitReader struct {
	data      []byte
	position  int
	bitOffset uint
	eof       bool
}

func newBitReader(data []byte) *bitReader {
	return &bitReader{data, 0, 0, false}
}

func (r *bitReader) EOF() bool {
	return r.eof
}

func (r *bitReader) Read1() uint32 {
	if r.position >= len(r.data) {
		r.eof = true
		return 0
	}
	var result uint32
	if r.data[r.position]&(1<<r.bitOffset) != 0 {
		result = 1
	}
	if r.bitOffset < 7 {
		r.bitOffset++
	} else {
		r.bitOffset = 0
		r.position++
	}
	return result
}

func (r *bitReader) Read8(n uint) uint8 {
	if n > 8 {
		panic("invalid argument")
	}
	var result uint8
	var written uint
	size := n
	for n > 0 {
		if r.position >= len(r.data) {
			r.eof = true
			return 0
		}
		result |= uint8(r.data[r.position]>>r.bitOffset) << written
		written += 8 - r.bitOffset
		if n < 8-r.bitOffset {
			r.bitOffset += n
			break
		}
		n -= 8 - r.bitOffset
		r.bitOffset = 0
		r.position++
	}
	return result &^ (0xFF << size)
}

func (r *bitReader) Read16(n uint) uint16 {
	if n > 16 {
		panic("invalid argument")
	}
	var result uint16
	var written uint
	size := n
	for n > 0 {
		if r.position >= len(r.data) {
			r.eof = true
			return 0
		}
		result |= uint16(r.data[r.position]>>r.bitOffset) << written
		written += 8 - r.bitOffset
		if n < 8-r.bitOffset {
			r.bitOffset += n
			break
		}
		n -= 8 - r.bitOffset
		r.bitOffset = 0
		r.position++
	}
	return result &^ (0xFFFF << size)
}

func (r *bitReader) Read32(n uint) uint32 {
	if n > 32 {
		panic("invalid argument")
	}
	var result uint32
	var written uint
	size := n
	for n > 0 {
		if r.position >= len(r.data) {
			r.eof = true
			return 0
		}
		result |= uint32(r.data[r.position]>>r.bitOffset) << written
		written += 8 - r.bitOffset
		if n < 8-r.bitOffset {
			r.bitOffset += n
			break
		}
		n -= 8 - r.bitOffset
		r.bitOffset = 0
		r.position++
	}
	return result &^ (0xFFFFFFFF << size)
}

func (r *bitReader) Read64(n uint) uint64 {
	if n > 64 {
		panic("invalid argument")
	}
	var result uint64
	var written uint
	size := n
	for n > 0 {
		if r.position >= len(r.data) {
			r.eof = true
			return 0
		}
		result |= uint64(r.data[r.position]>>r.bitOffset) << written
		written += 8 - r.bitOffset
		if n < 8-r.bitOffset {
			r.bitOffset += n
			break
		}
		n -= 8 - r.bitOffset
		r.bitOffset = 0
		r.position++
	}
	return result &^ (0xFFFFFFFFFFFFFFFF << size)
}

func (r *bitReader) ReadBool() bool {
	return r.Read8(1) == 1
}
