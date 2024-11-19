package main

import "testing"

func runSample(t *testing.T, m int, x []int, expect int) {
	sum, _ := solve(m, x)

	if expect != sum {
		t.Fatalf("Sample expect %d, but got %d", expect, sum)
	}

}

func TestSample1(t *testing.T) {
	m := 6
	x := []int{1, 5}
	expect := 8
	runSample(t, m, x, expect)
}

func TestSample2(t *testing.T) {
	m := 5
	x := []int{0, 3, 1}
	expect := 7
	runSample(t, m, x, expect)
}
