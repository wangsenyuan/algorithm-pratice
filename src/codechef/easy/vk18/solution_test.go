package main

import "testing"

func TestSamples(t *testing.T) {
	samples := []struct {
		n int
		e int
	}{
		{1, 2},
		{2, 12},
		{3, 36},
		{4, 80},
	}
	for _, sample := range samples {
		r := solve(sample.n)
		if r != sample.e {
			t.Errorf("sample %d should give %d, but got %d", sample.n, sample.e, r)
		}
	}
}
