package vorbis

type huffmanTable struct {
	table   []uint32
	visited []uint32
}

func newHuffmanTable(length uint32) *huffmanTable {
	return &huffmanTable{
		table:   make([]uint32, length),
		visited: make([]uint32, length),
	}
}

func (t huffmanTable) Lookup(r *bitReader) uint32 {
	i := uint32(0)
	for i&0x80000000 == 0 {
		i = t.table[i*2+r.Read1()]
	}
	return i & 0x7FFFFFFF
}

func (t huffmanTable) Put(entry uint32, length uint8) {
	t.put(0, entry, length-1)
}

func (t huffmanTable) put(index, entry uint32, length uint8) bool {
	if length < 32 && t.visited[index]&(1<<length) != 0 {
		return false
	}
	if length == 0 {
		if t.table[index*2] == 0 {
			t.table[index*2] = entry | 0x80000000
			return true
		}
		if t.table[index*2+1] == 0 {
			t.table[index*2+1] = entry | 0x80000000
			return true
		}
		t.visited[index] |= 1 << length
		return false
	}
	if t.table[index*2]&0x80000000 == 0 {
		if t.table[index*2] == 0 {
			t.table[index*2] = t.findEmpty(index + 1)
		}
		if t.put(t.table[index*2], entry, length-1) {
			return true
		}
	}
	if t.table[index*2+1]&0x80000000 == 0 {
		if t.table[index*2+1] == 0 {
			t.table[index*2+1] = t.findEmpty(index + 1)
		}
		if t.put(t.table[index*2+1], entry, length-1) {
			return true
		}
	}
	t.visited[index] |= 1 << length
	return false
}

func (t huffmanTable) findEmpty(index uint32) uint32 {
	for t.table[index*2] != 0 {
		index++
	}
	return index
}
