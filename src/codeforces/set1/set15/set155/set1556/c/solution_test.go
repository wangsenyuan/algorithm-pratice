package main

import "testing"

func runSample(t *testing.T, c []int, expect int) {
	res := solve(c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := []int{4, 1, 2, 3, 1}
	// (((()(()))(
	// ((()
	expect := 5
	runSample(t, c, expect)
}

func TestSample2(t *testing.T) {
	c := []int{1, 3, 2, 1, 2, 4}
	// (((()(()))(
	expect := 6
	runSample(t, c, expect)
}

func TestSample3(t *testing.T) {
	c := []int{1, 1, 1, 1, 2, 2}
	// (((()(()))(
	expect := 7
	runSample(t, c, expect)
}

func TestSample4(t *testing.T) {
	c := []int{60, 10, 62, 57}
	// (((()(()))(
	expect := 67
	runSample(t, c, expect)
}

func TestSample5(t *testing.T) {
	c := []int{18, 28, 35, 5, 35, 16, 41, 23, 21, 18, 41, 11, 28, 12, 39, 8, 14, 41, 26, 43, 14, 13, 36, 2, 20, 14, 15, 6, 6, 18, 40, 44, 12, 10, 12, 10, 33, 13, 14, 49, 48, 11, 48, 31, 50, 47, 35, 25, 10, 49}
	expect := 552
	runSample(t, c, expect)
}
