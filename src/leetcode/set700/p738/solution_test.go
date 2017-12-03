package p738

import "testing"

func TestSamples(t *testing.T) {
	samples := []struct {
		n int
		e int
	}{
		{10, 9},
		{1234, 1234},
		{332, 299},
		{31, 29},
		{101, 99},
	}

	for _, sample := range samples {
		res := monotoneIncreasingDigits(sample.n)
		if res != sample.e {
			t.Errorf("monotone increasing digits of %d should be %d, but got %d", sample.n, sample.e, res)
		}
	}
}
