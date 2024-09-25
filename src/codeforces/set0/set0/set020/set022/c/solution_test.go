package main

import "testing"

func runSample(t *testing.T, n int, m int, v int, expect bool) {
	res := solve(n, m, v)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	if len(res) != m {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 6, 3, true)
}


func TestSample2(t *testing.T) {
	runSample(t, 10, 26, 1, true)
}