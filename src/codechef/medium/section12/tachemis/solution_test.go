package main

import "testing"

func runSample(t *testing.T, n int, S []Pair, expect int64) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	S := make([]Pair, 4)
	S[1] = Pair{'A', 2}
	S[2] = Pair{'B', 2}
	var expect int64 = 6
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	S := make([]Pair, 5)
	S[1] = Pair{'A', 2}
	S[2] = Pair{'B', 1}
	S[3] = Pair{'A', 2}
	var expect int64 = 9
	runSample(t, n, S, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	S := make([]Pair, 3)
	S[1] = Pair{'A', 11}
	var expect int64 = 66
	runSample(t, n, S, expect)
}
