package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := "asdf"
	k := 5
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "aaaaa"
	k := 6
	expect := 15
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "aaaaa"
	k := 7
	expect := -1
	runSample(t, s, k, expect)
}
