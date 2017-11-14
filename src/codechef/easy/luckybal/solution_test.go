package main

import "testing"

func TestSample(t *testing.T) {
	tests := []struct {
		str    string
		expect int64
	}{
		{"47", 2},
		{"74", 2},
		{"477", 3},
		{"4747477", 23},
	}

	for _, test := range tests {
		res := solve(test.str)
		if res != test.expect {
			t.Errorf("solve(%s) should return %d, but is %d", test.str, test.expect, res)
		}
	}
}
