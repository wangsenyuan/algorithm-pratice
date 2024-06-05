package main

import "testing"

func runSample(t *testing.T, x int, c []int, expect int) {
	res := solve(x, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 3
	c := []int{2, 2, 2}
	expect := 2
	runSample(t, x, c, expect)
}

func TestSample2(t *testing.T) {
	x := 5
	c := []int{2, 2, 8, 2, 6, 8}
	expect := 4
	runSample(t, x, c, expect)
}

func TestSample3(t *testing.T) {
	x := 4
	c := []int{4, 10, 3, 8, 6, 10}
	expect := 3
	runSample(t, x, c, expect)
}

func TestSample4(t *testing.T) {
	x := 1
	c := []int{1, 1}
	expect := 1
	runSample(t, x, c, expect)
}

func TestSample5(t *testing.T) {
	x := 1
	c := []int{4, 1, 3, 1}
	expect := 2
	runSample(t, x, c, expect)
}

func TestSample6(t *testing.T) {
	x := 2
	c := []int{1, 3, 4, 3}
	expect := 1
	runSample(t, x, c, expect)
}
