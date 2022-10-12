package main

import "testing"

func runSample(t *testing.T, A []int, m int, n int, expect bool) {
	res := solve(A, m, n)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{12, 9, 8}
	m, n := 4, 6
	expect := true
	runSample(t, A, m, n, expect)
}

func TestSample2(t *testing.T) {
	A := []int{8, 8}
	m, n := 3, 3
	expect := false
	runSample(t, A, m, n, expect)
}

func TestSample3(t *testing.T) {
	A := []int{9, 5}
	m, n := 3, 3
	expect := true
	runSample(t, A, m, n, expect)
}

func TestSample4(t *testing.T) {
	A := []int{10, 11}
	m, n := 4, 5
	expect := true
	runSample(t, A, m, n, expect)
}

func TestSample5(t *testing.T) {
	A := []int{9, 11}
	m, n := 5, 4
	expect := false
	runSample(t, A, m, n, expect)
}

func TestSample6(t *testing.T) {
	A := []int{11, 45, 14}
	m, n := 10, 10
	expect := false
	runSample(t, A, m, n, expect)
}
