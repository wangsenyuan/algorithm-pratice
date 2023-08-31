package main

import "testing"

func runSample(t *testing.T, A []int, m int, B []int, expect bool) {
	res := solve(A, m, B)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 2, 4, 2}
	m := 2
	B := []int{1, 4, 4, 2}
	expect := true
	runSample(t, A, m, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 2, 8, 2, 2}
	m := 2
	B := []int{1, 16}
	expect := true
	runSample(t, A, m, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 3, 3, 3, 3, 3, 3, 3}
	m := 3
	B := []int{6, 6, 6, 6}
	expect := false
	runSample(t, A, m, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{3, 9, 6, 3, 12, 12, 36, 12}
	m := 3
	B := []int{9, 3, 2, 2, 2, 3, 4, 12, 4, 12, 4, 12, 4, 12, 4, 4}
	expect := true
	runSample(t, A, m, B, expect)
}

func TestSample5(t *testing.T) {
	A := []int{3, 9, 6, 3, 12, 12, 36, 12}
	m := 3
	B := []int{12, 2, 4, 3, 4, 12, 56}
	expect := false
	runSample(t, A, m, B, expect)
}
