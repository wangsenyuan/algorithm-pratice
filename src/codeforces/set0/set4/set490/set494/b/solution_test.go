package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababa"
	x := "aba"
	// [0, 2], [2, 4]
	// 只有一个区间时, []
	expect := 5
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "ddd"
	x := "d"
	expect := 12
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "welcometoroundtwohundredandeightytwo"
	x := "d"
	expect := 274201
	runSample(t, s, x, expect)
}
