package main

import "testing"

func runSample(t *testing.T, a []int, expect string) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5}
	expect := "Alice"
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 4, 5}
	expect := "Alice"
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 8, 2, 1, 10, 9}
	expect := "Bob"
	runSample(t, a, expect)
}