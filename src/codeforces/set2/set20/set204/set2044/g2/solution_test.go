package main

import "testing"

func runSample(t *testing.T, r []int, expect int) {
	res := solve(r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	r := []int{2, 1}
	expect := 2
	runSample(t, r, expect)
}

func TestSample2(t *testing.T) {
	r := []int{2, 3, 4, 5, 1}
	expect := 2
	runSample(t, r, expect)
}

func TestSample3(t *testing.T) {
	r := []int{2, 1, 4, 2, 3}
	expect := 5
	runSample(t, r, expect)
}

func TestSample4(t *testing.T) {
	r := []int{4, 1, 1, 5, 4}
	expect := 5
	runSample(t, r, expect)
}

func TestSample5(t *testing.T) {
	r := []int{4, 3, 9, 1, 6, 7, 9, 10, 10, 3}
	expect := 5
	runSample(t, r, expect)
}

func TestSample6(t *testing.T) {
	r := []int{29, 1, 2, 1, 2, 4, 6, 7, 4, 1, 3, 4, 6, 13, 11, 1, 9, 10, 3, 7, 1, 9, 20, 11, 22, 3, 5, 18, 9, 23, 23, 30, 21, 13, 1, 32, 11, 29, 17, 11, 12, 17, 7, 42, 34, 39, 21, 12, 13, 30, 50, 9, 39, 16, 25, 50, 5, 38, 15, 46}
	expect := 20
	runSample(t, r, expect)
}

func TestSample7(t *testing.T) {
	r := []int{11, 1, 1, 3, 3, 1, 1, 5, 7, 8, 4, 6, 12}
	expect := 5
	runSample(t, r, expect)
}
