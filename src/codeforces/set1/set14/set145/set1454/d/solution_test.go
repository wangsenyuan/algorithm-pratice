package main

import "testing"

func runSample(t *testing.T, n int64, expect int) {
	res := solve(n)
	if len(res) != expect {
		t.Fatalf("Sample expect len %d result, but got %v", expect, res)
	}

	var prod int64 = 1
	for i := 0; i < len(res); i++ {
		prod *= res[i]
		if i > 0 && res[i]%res[i-1] != 0 {
			t.Fatalf("Sample result %v, not correct at pos %d", res, i)
		}
	}
	if prod != n {
		t.Fatalf("Sample result %v, get wrong product %d, not %d", res, prod, n)
	}
}

func TestSample1(t *testing.T) {
	var n int64 = 2
	expect := 1
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	var n int64 = 360
	expect := 3
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	var n int64 = 4999999937
	expect := 1
	runSample(t, n, expect)
}
func TestSample4(t *testing.T) {
	var n int64 = 4998207083
	expect := 1
	runSample(t, n, expect)
}
