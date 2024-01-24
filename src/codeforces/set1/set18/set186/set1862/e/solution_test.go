package main

import "testing"

func runSample(t *testing.T, m int, d int, a []int, expect int) {
	res := solve(m, d, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, d := 2, 2
	a := []int{3, 2, 5, 4, 6}
	expect := 2
	runSample(t, m, d, a, expect)
}
func TestSample2(t *testing.T) {
	m, d := 3, 2
	a := []int{1, 1, 1, 1}
	expect := 0
	runSample(t, m, d, a, expect)
}

func TestSample3(t *testing.T) {
	m, d := 6, 6
	a := []int{-82, 45, 1, -77, 39, 11}
	expect := 60
	runSample(t, m, d, a, expect)
}

func TestSample4(t *testing.T) {
	m, d := 2, 2
	a := []int{3, 2, 5, 4, 8}
	expect := 3
	runSample(t, m, d, a, expect)
}

func TestSample5(t *testing.T) {
	m, d := 3, 2
	a := []int{-8, 8, -2, -1, 9, 0}
	expect := 7
	runSample(t, m, d, a, expect)
}
