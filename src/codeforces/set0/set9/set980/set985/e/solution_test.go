package main

import "testing"

func runSample(t *testing.T, k int, d int, a []int, expect bool) {
	res := solve(k, d, a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, d := 3, 10
	a := []int{7, 2, 7, 7, 4, 2}
	expect := true
	runSample(t, k, d, a, expect)
}

func TestSample2(t *testing.T) {
	k, d := 2, 3
	a := []int{4, 5, 3, 13, 4, 10}
	expect := true
	runSample(t, k, d, a, expect)
}

func TestSample3(t *testing.T) {
	k, d := 2, 5
	a := []int{10, 16, 22}
	expect := false
	runSample(t, k, d, a, expect)
}
func TestSample4(t *testing.T) {
	k, d := 2, 76
	a := []int{44, 5, 93}
	expect := false
	runSample(t, k, d, a, expect)
}
