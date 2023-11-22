package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := false
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 1, 4, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 3, 2, 4, 6, 5, 7}
	expect := true
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 3, 6, 5, 4, 1}
	expect := false
	runSample(t, a, expect)
}

// 1 2 3 5 4 7 6
func TestSample5(t *testing.T) {
	a := []int{1, 2, 3, 5, 4, 7, 6}
	expect := true
	runSample(t, a, expect)
}
