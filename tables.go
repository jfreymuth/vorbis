package vorbis

import (
	"math/bits"
	"sync"
)

// blockTables holds the window and the IMDCT lookup for one blocksize. Both
// depend on nothing but the blocksize, so one set serves every decoder in the
// process: they are built on first use and never written again, which lets
// decoders on different goroutines read them without a lock.
type blockTables struct {
	once   sync.Once
	window []float32
	imdct  imdctLookup
}

// tables is indexed by the blocksize's exponent, as the identification header
// stores it in a nibble. Vorbis allows blocksizes 2^6 to 2^13.
var tables [16]blockTables

// tablesFor returns the tables for a blocksize, building them the first time.
func tablesFor(blocksize int) *blockTables {
	t := &tables[bits.TrailingZeros(uint(blocksize))]
	t.once.Do(func() {
		t.window = makeWindow(blocksize)
		generateIMDCTLookup(blocksize, &t.imdct)
	})
	return t
}
