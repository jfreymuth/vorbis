package vorbis

import "testing"

// TestTablesShared checks that decoders for the same blocksizes read the same
// window and IMDCT tables, and that those are what a fresh build would give.
func TestTablesShared(t *testing.T) {
	data, err := readTestFile()
	if err != nil {
		t.Fatal(err)
	}
	var a, b Decoder
	for _, header := range data.Headers {
		if err := a.ReadHeader(header); err != nil {
			t.Fatal(err)
		}
		if err := b.ReadHeader(header); err != nil {
			t.Fatal(err)
		}
	}
	for i := range a.blocksize {
		if &a.windows[i][0] != &b.windows[i][0] || a.lookup[i] != b.lookup[i] {
			t.Errorf("blocksize %d: decoders do not share tables", a.blocksize[i])
		}
		window := makeWindow(a.blocksize[i])
		if len(window) != len(a.windows[i]) {
			t.Fatalf("blocksize %d: window has %d entries, want %d", a.blocksize[i], len(a.windows[i]), len(window))
		}
		for j := range window {
			if window[j] != a.windows[i][j] {
				t.Fatalf("blocksize %d: window[%d] = %g, want %g", a.blocksize[i], j, a.windows[i][j], window[j])
			}
		}
		var lookup imdctLookup
		generateIMDCTLookup(a.blocksize[i], &lookup)
		for name, pair := range map[string][2][]float32{"A": {lookup.A, a.lookup[i].A}, "B": {lookup.B, a.lookup[i].B}, "C": {lookup.C, a.lookup[i].C}} {
			if len(pair[0]) != len(pair[1]) {
				t.Fatalf("blocksize %d: lookup %s has %d entries, want %d", a.blocksize[i], name, len(pair[1]), len(pair[0]))
			}
			for j := range pair[0] {
				if pair[0][j] != pair[1][j] {
					t.Fatalf("blocksize %d: lookup %s[%d] = %g, want %g", a.blocksize[i], name, j, pair[1][j], pair[0][j])
				}
			}
		}
	}
}
