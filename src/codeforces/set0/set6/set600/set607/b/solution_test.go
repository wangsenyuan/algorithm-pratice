package main

import "testing"

func runSample(t *testing.T, c []int, expect int) {
	res := solve(c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := []int{1, 2, 1}
	expect := 1
	runSample(t, c, expect)
}

func TestSample2(t *testing.T) {
	c := []int{1, 2, 3}
	expect := 3
	runSample(t, c, expect)
}

func TestSample3(t *testing.T) {
	c := []int{1, 4, 4, 2, 3, 2, 1}
	expect := 2
	runSample(t, c, expect)
}

func TestSample4(t *testing.T) {
	c := []int{1, 1}
	expect := 1
	runSample(t, c, expect)
}
