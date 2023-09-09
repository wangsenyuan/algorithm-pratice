package main

import "testing"

func runSample(t *testing.T, x []int, y []int, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{8, 3, 9, 2, 4, 5}
	y := []int{5, 3, 1, 4, 5, 10}
	expect := 2
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x := []int{1, 2, 3, 4}
	y := []int{1, 1, 2, 2}
	expect := 0
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x := []int{2, 3, 7}
	y := []int{1, 3, 10}
	expect := 1
	runSample(t, x, y, expect)
}

func TestSample4(t *testing.T) {
	x := []int{2, 3, 6, 9, 5, 7}
	y := []int{3, 2, 7, 10, 6, 10}
	expect := 3
	runSample(t, x, y, expect)
}

func TestSample5(t *testing.T) {
	x := []int{5, 4, 2, 1, 8, 100}
	y := []int{1, 1, 1, 1, 1, 200}
	expect := 1
	runSample(t, x, y, expect)
}

func TestSample6(t *testing.T) {
	x := []int{1, 4, 1, 2, 4, 2}
	y := []int{1, 3, 3, 2, 3, 4}
	expect := 3
	runSample(t, x, y, expect)
}
