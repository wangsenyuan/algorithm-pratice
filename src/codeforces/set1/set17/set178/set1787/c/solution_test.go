package main

import "testing"

func runSample(t *testing.T, a []int, s int, expect int) {
	res := solve(a, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := 0
	a := []int{2, 0, 1, 3, 4}
	expect := 0
	runSample(t, a, s, expect)
}

func TestSample2(t *testing.T) {
	s := 1
	a := []int{5, 3, 4, 3, 5}
	expect := 18
	runSample(t, a, s, expect)
}

func TestSample3(t *testing.T) {
	s := 2
	a := []int{7, 6, 5, 4, 3, 2, 1}
	expect := 32
	runSample(t, a, s, expect)
}

func TestSample4(t *testing.T) {
	s := 1
	a := []int{1, 2, 3, 4, 5}
	expect := 11
	runSample(t, a, s, expect)
}

func TestSample5(t *testing.T) {
	s := 2
	a := []int{1, 2, 3, 4, 5}
	expect := 14
	runSample(t, a, s, expect)
}

func TestSample6(t *testing.T) {
	s := 0
	a := []int{0, 1, 1, 1}
	expect := 0
	runSample(t, a, s, expect)
}

func TestSample7(t *testing.T) {
	s := 5
	a := []int{4, 3, 5, 6, 4}
	expect := 16
	runSample(t, a, s, expect)
}

func TestSample8(t *testing.T) {
	s := 8139
	a := []int{7976, 129785, 12984, 78561, 173685, 15480}
	expect := 2700826806
	runSample(t, a, s, expect)
}

func TestSample9(t *testing.T) {
	s := 99999
	a := []int{200000, 200000, 200000}
	expect := 40000000000
	runSample(t, a, s, expect)
}
