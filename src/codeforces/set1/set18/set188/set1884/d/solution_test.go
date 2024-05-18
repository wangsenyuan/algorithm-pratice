package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 4, 4, 4}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 3, 4, 4}
	expect := 3
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{6, 8, 9, 4, 6, 8, 9, 4, 9}
	expect := 26
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{7, 7, 4, 4, 9, 9, 6, 2, 9}
	expect := 26
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{10, 18, 18, 15, 14, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 17, 13, 11}
	expect := 124
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{12, 19, 19, 18, 18, 12, 2, 18, 19, 12, 12, 3, 12, 12, 12, 18, 19, 16, 18, 19, 12}
	expect := 82
	runSample(t, a, expect)
}
