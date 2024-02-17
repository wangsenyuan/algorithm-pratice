package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 2}
	expect := 12
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{39, 68, 31, 80}
	expect := 1337
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{313539461, 779847196, 221612534, 488613315, 633203958, 394620685, 761188160}
	expect := 257421502
	runSample(t, a, expect)
}
