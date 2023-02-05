package main

import "testing"

func runSample(t *testing.T, A []int, m int, expect string) {
	res := solve(A, m)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 3, 1}
	m := 5
	expect := "ABABA"
	runSample(t, A, m, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2}
	m := 5
	expect := "BABBB"
	runSample(t, A, m, expect)
}
