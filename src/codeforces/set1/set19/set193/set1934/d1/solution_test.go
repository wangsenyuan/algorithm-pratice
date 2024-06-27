package main

import "testing"

func runSample(t *testing.T, n int, m int, expect bool) {
	res := solve(n, m)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	if res[0] != n || res[len(res)-1] != m {
		t.Fatalf("Sample result %v, not correct", res)
	}

	for i := 1; i < len(res); i++ {
		x := res[i-1]
		y := res[i]
		if y > x || x^y > x {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 3, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 481885160128643072, 45035996273704960, true)
}
