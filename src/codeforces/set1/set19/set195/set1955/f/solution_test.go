package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1, 0}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 0, 1, 2}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 2, 2, 0}
	expect := 3
	runSample(t, a, expect)
}
func TestSample4(t *testing.T) {
	a := []int{3, 3, 2, 0}
	expect := 3
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{0, 9, 9, 9}
	expect := 12
	runSample(t, a, expect)
}