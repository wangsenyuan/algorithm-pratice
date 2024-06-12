package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{20, 6, 12, 3, 48, 36}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{12, 6, 3, 4}
	expect := false
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{10, 12, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{32, 16, 8, 4, 2}
	expect := false
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{100, 50, 2, 10, 20}
	expect := true
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, 6, 11, 12, 6, 12, 3, 6}
	expect := true
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{2, 4, 8, 1}
	expect := true
	runSample(t, a, expect)
}
