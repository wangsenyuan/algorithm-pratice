package main

import "testing"

func runSample(t *testing.T, b int, x int, c []int, expect int) {
	res := solve(b, x, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	b, x := 1, 0
	c := []int{1, 2, 3}
	expect := 4
	runSample(t, b, x, c, expect)
}

func TestSample2(t *testing.T) {
	b, x := 5, 10
	c := []int{2, 5, 3}
	expect := 40
	runSample(t, b, x, c, expect)
}

func TestSample3(t *testing.T) {
	b, x := 3, 3
	c := []int{3, 2, 1, 2}
	expect := 9
	runSample(t, b, x, c, expect)
}

func TestSample4(t *testing.T) {
	b, x := 1, 0
	c := []int{4, 1, 4, 2}
	expect := 13
	runSample(t, b, x, c, expect)
}

func TestSample5(t *testing.T) {
	b, x := 1, 10
	c := []int{4, 1, 4, 2}
	expect := 0
	runSample(t, b, x, c, expect)
}
