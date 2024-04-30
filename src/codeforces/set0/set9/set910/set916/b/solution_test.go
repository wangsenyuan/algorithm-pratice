package main

import (
	"testing"
)

func runSample(t *testing.T, n, k int, expect int) {
	res := solve(n, k)

	if len(res) == 0 {
		if expect == -100 {
			return
		}
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if res[0] != expect || len(res) != k {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 23, 5
	expect := 3
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 13, 2
	runSample(t, n, k, -100)
}

func TestSample3(t *testing.T) {
	n, k := 1, 2
	runSample(t, n, k, -1)
}

func TestSample4(t *testing.T) {
	n, k := 1000000000000000000, 100000
	expect := 44
	runSample(t, n, k, expect)
}
