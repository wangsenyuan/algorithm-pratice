package main

import "testing"

func runSample(t *testing.T, s string, p []int, expect int) {
	res := solve(s, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababa"
	p := []int{3, 4, 5, 2, 1}
	expect := 1
	runSample(t, s, p, expect)
}

func TestSample2(t *testing.T) {
	s := "ababa"
	p := []int{2, 1, 4, 5, 3}
	expect := 6
	runSample(t, s, p, expect)
}

func TestSample3(t *testing.T) {
	s := "codeforces"
	p := []int{8, 6, 1, 7, 5, 2, 9, 3, 10, 4}
	expect := 12
	runSample(t, s, p, expect)
}
