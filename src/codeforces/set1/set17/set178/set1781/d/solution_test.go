package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 6, 13, 22, 97}
	expect := 5
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{101}
	expect := 1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 5, 10, 17, 26}
	expect := 2
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{18242509, 425905756, 528992469, 617747484}
	expect := 3
	runSample(t, a, expect)
}
