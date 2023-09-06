package main

import "testing"

func runSample(t *testing.T, a, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 0, 0, 3, 3}
	b := []int{2, 3, 2, 1, 0}
	expect := 2
	runSample(t, a, b, expect)
}
func TestSample2(t *testing.T) {
	a := []int{1, 1, 1}
	b := []int{0, 0, 3}
	expect := 0
	runSample(t, a, b, expect)
}
func TestSample3(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	b := []int{0, 1, 2, 3, 4, 5, 6, 7}
	expect := 7
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{28, 14, 12, 27, 10, 8, 27, 27}
	b := []int{5, 23, 17, 2, 21, 19, 6, 22}
	expect := 12
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := []int{28, 12, 15, 14, 2, 28, 10, 10, 4, 25, 19, 5, 15}
	b := []int{23, 11, 16, 17, 9, 23, 1, 9, 19, 22, 12, 18, 8}
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample6(t *testing.T) {
	a := []int{27, 30, 0}
	b := []int{0, 5, 9}
	expect := 9
	runSample(t, a, b, expect)
}
