package main

import "testing"

func TestSample(t *testing.T) {
	sample := []struct {
		n      int
		m      int
		expect int64
	}{
		{2, 2, 12},
		{2, 3, 26},
		{4, 5, 312},
	}

	for _, each := range sample {
		res := solve(each.n, each.m)
		if res != each.expect {
			t.Errorf("sample %d %d should give %d, but got %d", each.n, each.m, each.expect, res)
		}
	}
}
