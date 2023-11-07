package main

import "testing"

func runSample(t *testing.T, a []int, m int, expect int) {
	res := solve(len(a), a, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{100, 100, 200, 1}
	m := 401
	expect := 1
	runSample(t, a, m, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	m := 2
	expect := 2
	runSample(t, a, m, expect)
}
func TestSample3(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	m := 0
	expect := 6
	runSample(t, a, m, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 2, 1}
	m := 4
	expect := 1
	runSample(t, a, m, expect)
}

func TestSample5(t *testing.T) {
	a := []int{0, 1, 1, 1}
	m := 0
	expect := 4
	runSample(t, a, m, expect)
}
