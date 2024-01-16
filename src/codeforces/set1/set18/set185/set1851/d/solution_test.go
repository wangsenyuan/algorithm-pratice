package main

import "testing"

func runSample(t *testing.T, sums []int, expect bool) {
	res := solve(sums)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	sums := []int{6, 8, 12, 15}
	expect := true
	runSample(t, sums, expect)
}

func TestSample2(t *testing.T) {
	sums := []int{1, 6, 8, 15}
	expect := true
	runSample(t, sums, expect)
}

func TestSample3(t *testing.T) {
	sums := []int{1, 2, 100}
	expect := false
	runSample(t, sums, expect)
}

func TestSample4(t *testing.T) {
	sums := []int{1, 3, 6}
	expect := true
	runSample(t, sums, expect)
}

func TestSample5(t *testing.T) {
	sums := []int{2}
	expect := true
	runSample(t, sums, expect)
}

func TestSample6(t *testing.T) {
	sums := []int{1, 2}
	expect := false
	runSample(t, sums, expect)
}

func TestSample7(t *testing.T) {
	sums := []int{9, 11, 12, 20, 25, 28, 30, 33}
	expect := false
	runSample(t, sums, expect)
}
