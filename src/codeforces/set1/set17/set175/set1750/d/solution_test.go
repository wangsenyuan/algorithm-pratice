package main

import "testing"

func runSample(t *testing.T, m int, a []int, expect int) {
	res := solve(m, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 5
	a := []int{4, 2, 1}
	// 2, 1
	// 2, 3
	// 2, 5
	expect := 3
	runSample(t, m, a, expect)
}

func TestSample2(t *testing.T) {
	m := 1
	a := []int{1, 1}
	// 2, 1
	// 2, 3
	// 2, 5
	expect := 1
	runSample(t, m, a, expect)
}
func TestSample3(t *testing.T) {
	m := 5
	a := []int{2, 3, 5, 2, 3}
	// 2, 1
	// 2, 3
	// 2, 5
	expect := 0
	runSample(t, m, a, expect)
}

func TestSample4(t *testing.T) {
	m := 1000000000
	a := []int{60, 30, 1, 1}
	// 2, 1
	// 2, 3
	// 2, 5
	expect := 595458194
	runSample(t, m, a, expect)
}

func TestSample5(t *testing.T) {
	m := 1000000000
	a := []int{1000000000, 2}
	// 2, 1
	// 2, 3
	// 2, 5
	expect := 200000000
	runSample(t, m, a, expect)
}
