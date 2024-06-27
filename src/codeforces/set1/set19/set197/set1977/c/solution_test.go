package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	a := []int{1, 2, 4, 8, 16}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 2, 10, 20, 60, 1}
	expect := 4
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 3, 4, 6, 12, 100003, 1200036}
	expect := 4
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 42, 7, 3, 6, 7, 7, 1, 6}
	expect := 5
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{4, 99, 57, 179, 10203, 2, 11, 40812}
	expect := 8
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{39083, 19, 2299, 11, 2057, 17}
	expect := 3
	runSample(t, a, expect)
}
