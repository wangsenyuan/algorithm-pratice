package main

import "testing"

func runSample(t *testing.T, C []int, expect bool) {
	res := solve(C)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{1}
	expect := true
	runSample(t, C, expect)
}

func TestSample2(t *testing.T) {
	C := []int{1, 1}
	expect := true
	runSample(t, C, expect)
}

func TestSample3(t *testing.T) {
	C := []int{1, 2}
	expect := false
	runSample(t, C, expect)
}

func TestSample4(t *testing.T) {
	C := []int{3, 1, 3}
	expect := true
	runSample(t, C, expect)
}

func TestSample5(t *testing.T) {
	C := []int{26, 8, 3, 4, 13, 34}
	expect := true
	runSample(t, C, expect)
}

func TestSample6(t *testing.T) {
	C := []int{7, 5}
	expect := false
	runSample(t, C, expect)
}
