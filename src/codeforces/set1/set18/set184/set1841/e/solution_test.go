package main

import "testing"

func runSample(t *testing.T, a []int, m int, expect int) {
	res := solve(a, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 0, 0}
	m := 9
	expect := 6
	runSample(t, a, m, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 0, 3, 1}
	m := 5
	expect := 3
	runSample(t, a, m, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 0, 3, 1}
	m := 6
	expect := 4
	runSample(t, a, m, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0, 2, 2, 1, 5, 10, 3, 4, 1, 1}
	m := 20
	expect := 16
	runSample(t, a, m, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1}
	m := 1
	expect := 0
	runSample(t, a, m, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 0, 3, 1}
	m := 10
	expect := 4
	runSample(t, a, m, expect)
}
