package main

import (
	"testing"
)

func runSample(t *testing.T, a []int, x int, expect int) {
	res := solve(a, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	x := 2
	expect := 2
	runSample(t, a, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1}
	x := 2
	expect := 2
	runSample(t, a, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 3}
	x := 2
	expect := 1
	runSample(t, a, x, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0, 0}
	x := 3
	expect := 2
	runSample(t, a, x, expect)
}

func TestSample5(t *testing.T) {
	a := []int{0, 0, 1}
	x := 2
	expect := 3
	runSample(t, a, x, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, 3, 3, 7}
	x := 2
	expect := -1
	runSample(t, a, x, expect)
}
func TestSample7(t *testing.T) {
	a := []int{2, 3}
	x := 2
	expect := 1
	runSample(t, a, x, expect)
}

func TestSample8(t *testing.T) {
	a := []int{0, 1, 2, 2, 1}
	x := 0
	expect := 2
	runSample(t, a, x, expect)
}
