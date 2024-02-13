package main

import "testing"

func runSample(t *testing.T, s string, L []int, R []int, ops []int, expect string) {
	res := solve(s, L, R, ops)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcd"
	L := []int{1, 3}
	R := []int{2, 4}
	ops := []int{1, 3}
	expect := "badc"
	runSample(t, s, L, R, ops, expect)
}
