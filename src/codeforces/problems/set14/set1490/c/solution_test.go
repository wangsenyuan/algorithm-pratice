package main

import "testing"

func runSample(t *testing.T, x int64, expect bool) {
	res := solve(x)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := []int64{
		1,
		2,
		4,
		34,
		35,
		16,
		703657519796,
	}
	expect := []bool{
		false, true, false, false, true, true, true,
	}

	for i := 0; i < len(X); i++ {
		runSample(t, X[i], expect[i])
	}
}
