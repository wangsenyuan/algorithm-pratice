package main

import "testing"

func runSample(t *testing.T, jars []int, expect int) {
	res := solve(jars)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	jars := []int{1, 1, 1, 2, 2, 1, 2, 1, 2, 1, 1, 2}
	expect := 6
	runSample(t, jars, expect)
}

func TestSample2(t *testing.T) {
	jars := []int{1, 2, 1, 2}
	expect := 0
	runSample(t, jars, expect)
}
