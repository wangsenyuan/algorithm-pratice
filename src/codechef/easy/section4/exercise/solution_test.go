package main

import "testing"

func runSample(t *testing.T, n int, A []int, S string, expect bool) {
	res := solve(n, A, S)
	if res != expect {
		t.Errorf("Sample %d %v %s, expect %t, but got %t", n, A, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{7, 5, -1, -1, -1, 10}
	S := ">>=>>"
	expect := false
	runSample(t, n, A, S, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	A := []int{7, 5, -1, -1, -1, 10}
	S := ">>=<>"
	expect := true
	runSample(t, n, A, S, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	A := []int{0, -1}
	S := ">"
	expect := false
	runSample(t, n, A, S, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	A := []int{5, -1, -1, -1, 10}
	S := "<<<<"
	expect := true
	runSample(t, n, A, S, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	A := []int{5, -1, -1, -1, 8}
	S := "<<<<"
	expect := false
	runSample(t, n, A, S, expect)
}
