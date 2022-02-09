package main

import "testing"

func runSample(t *testing.T, n int, k int, s string, expect int) {
	res := solve(n, k, s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 5
	s := "380"
	expect := 0
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 9
	s := "380"
	expect := 3
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 4, 2
	s := "0123"
	expect := 1
	runSample(t, n, k, s, expect)
}

func TestSample4(t *testing.T) {
	n, k := 5, 13
	s := "78712"
	expect := 3
	runSample(t, n, k, s, expect)
}

func TestSample5(t *testing.T) {
	n, k := 6, 10
	s := "051827"
	expect := 2
	runSample(t, n, k, s, expect)
}


func TestSample6(t *testing.T) {
	n, k := 8, 25
	s := "37159725"
	expect := 5
	runSample(t, n, k, s, expect)
}