package main

import "testing"

func runSample(t *testing.T, m int, k int, a []int, expect int) {
	res := solve(m, k, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, k := 3, 2
	a := []int{3, 5, 1}
	expect := 1
	runSample(t, m, k, a, expect)
}

func TestSample2(t *testing.T) {
	m, k := 10, 3
	a := []int{12, 8, 18, 25, 1}
	expect := 0
	runSample(t, m, k, a, expect)
}

func TestSample3(t *testing.T) {
	m, k := 7, 2
	a := []int{7, 3, 4, 1, 6, 5, 2}
	expect := 6
	runSample(t, m, k, a, expect)
}

func TestSample4(t *testing.T) {
	m, k := 2, 2
	a := []int{1, 3}
	expect := 0
	runSample(t, m, k, a, expect)
}

func TestSample5(t *testing.T) {
	m, k := 60, 1
	a := []int{223061, 155789, 448455, 956209, 90420, 110807, 833270, 240866, 996739, 14579, 366906, 594384, 72757, 50161, 278465, 135449}
	expect := 16
	runSample(t, m, k, a, expect)
}
