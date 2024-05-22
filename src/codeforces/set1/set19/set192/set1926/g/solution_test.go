package main

import "testing"

func runSample(t *testing.T, n int, s string, a []int, expect int) {
	res := solve(n, a, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := []int{1, 1}
	s := "CSP"
	expect := 1
	runSample(t, n, s, a, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	a := []int{1, 2, 2}
	s := "PCSS"
	expect := 1
	runSample(t, n, s, a, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	a := []int{1, 2, 2}
	s := "PPSS"
	expect := 2
	runSample(t, n, s, a, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	a := []int{1, 1, 2, 1, 1}
	s := "CSPPCS"
	expect := 2
	runSample(t, n, s, a, expect)
}

func TestSample5(t *testing.T) {
	n := 7
	a := []int{1, 2, 3, 3, 4, 6}
	s := "CPSPPPP"
	expect := 3
	runSample(t, n, s, a, expect)
}

func TestSample6(t *testing.T) {
	n := 4
	a := []int{1, 2, 3}
	s := "CCPS"
	expect := 1
	runSample(t, n, s, a, expect)
}
