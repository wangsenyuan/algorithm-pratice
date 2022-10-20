package main

import "testing"

func runSample(t *testing.T, n int, X []int, P []int, m int, expect string) {
	res := solve(n, X, P, m)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	m := 6
	X := []int{1, 5, 3}
	P := []int{5, 5, 4}
	expect := "001"
	runSample(t, n, X, P, m, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	m := 3
	X := []int{1, 5}
	P := []int{3, 2}
	expect := "11"
	runSample(t, n, X, P, m, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	m := 5
	X := []int{1, 10}
	P := []int{6, 6}
	expect := "00"
	runSample(t, n, X, P, m, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	m := 12
	X := []int{4, 1, 12, 5, 9, 8}
	P := []int{5, 6, 5, 5, 7, 3}
	expect := "100110"
	runSample(t, n, X, P, m, expect)
}
