package main

import "testing"

func runSample(t *testing.T, A string, M int, expect int) {
	res := int(solve(A, M))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "00"
	m := 2
	expect := 4
	runSample(t, A, m, expect)
}

func TestSample2(t *testing.T) {
	A := "11"
	m := 4
	expect := 1
	runSample(t, A, m, expect)
}

func TestSample3(t *testing.T) {
	A := "101"
	m := 3
	expect := 2
	runSample(t, A, m, expect)
}
