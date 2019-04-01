package popcount

import "testing"

func testdata() []uint64 {
	var d []uint64
	for i := 0; i < 10000; i++ {
		d = append(d, uint64(i))
	}
	return d
}

func BenchmarkPopCount(b *testing.B) {
	d := testdata()
	for i := 0; i < b.N; i++ {
		for _, u := range d {
			PopCount(u)
		}
	}
}

func BenchmarkPopCountUsingLoop(b *testing.B) {
	d := testdata()
	for i := 0; i < b.N; i++ {
		for _, u := range d {
			PopCountUsingLoop(u)
		}
	}
}
