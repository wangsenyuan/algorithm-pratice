package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 5, 3, 2, 4}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 1, 3, 4, 5, 6, 7, 8}
	expect := 3
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 4, 4}
	expect := 2
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{4, 4, 1, 3, 5}
	expect := 0
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1}
	expect := 0
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{22, 16, 12, 20, 21, 19, 14, 20, 4, 22, 19, 11, 16, 6, 15, 2, 17, 10, 14, 9, 19, 21}
	expect := 3
	runSample(t, a, expect)
}

func TestSample8(t *testing.T) {
	a := []int{100, 2, 3, 78, 59, 86, 62, 8, 9, 86, 11, 19, 13, 61, 15, 50, 17, 25, 19, 91, 21, 44, 23, 82, 25, 56, 27, 89, 29, 58, 31, 86, 33, 75, 35, 60, 37, 4, 5, 65, 67, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	expect := 48
	runSample(t, a, expect)
}
