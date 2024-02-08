package main

import "testing"

func runSample(t *testing.T, l int, r int, expect bool) {
	res := solve(l, r)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	a, b := res[0], res[1]
	x := a + b
	if x < l || x > r || gcd(a, b) == 1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 9, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 3, false)
}
func TestSample3(t *testing.T) {
	runSample(t, 18, 19, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 777, 777, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 1791791, 1791791, false)
}
