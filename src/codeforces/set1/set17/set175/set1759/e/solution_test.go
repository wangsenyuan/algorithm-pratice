package main

import "testing"

func runSample(t *testing.T, a []int, h int, expect int) {
	res := solve(a, h)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := 1
	a := []int{2, 1, 8, 9}
	expect := 4
	runSample(t, a, h, expect)
}

func TestSample2(t *testing.T) {
	h := 3
	a := []int{6, 2, 60}
	expect := 3
	runSample(t, a, h, expect)
}

func TestSample3(t *testing.T) {
	h := 5
	a := []int{5, 1, 100, 5}
	expect := 3
	runSample(t, a, h, expect)
}

func TestSample4(t *testing.T) {
	h := 2
	a := []int{38, 6, 3}
	expect := 3
	runSample(t, a, h, expect)
}

func TestSample5(t *testing.T) {
	h := 1
	a := []int{12}
	expect := 0
	runSample(t, a, h, expect)
}

func TestSample6(t *testing.T) {
	h := 6
	a := []int{12, 12, 36, 100}
	expect := 4
	runSample(t, a, h, expect)
}

func TestSample7(t *testing.T) {
	h := 1
	a := []int{2, 1, 1, 15}
	expect := 4
	runSample(t, a, h, expect)
}

func TestSample8(t *testing.T) {
	h := 5
	a := []int{15, 1, 13}
	expect := 3
	runSample(t, a, h, expect)
}
