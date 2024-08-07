package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{1}
	expect := 0
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{1, 2}
	expect := 1
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{2, 3, 1}
	expect := 1
	runSample(t, p, expect)
}
func TestSample4(t *testing.T) {
	p := []int{2, 4, 1, 3, 5}
	expect := 3
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{8, 9, 7, 12, 1, 10, 6, 3, 2, 4, 11, 5}
	expect := 9
	runSample(t, p, expect)
}

func TestSample6(t *testing.T) {
	p := []int{1, 2, 4, 6, 8, 10, 12, 14, 3, 9, 15, 5, 7, 11, 13}
	expect := 3
	runSample(t, p, expect)
}
