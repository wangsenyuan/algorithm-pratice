package main

import "testing"

func runSample(t *testing.T, n int, m int, expect bool) {
	res := solve(n, m)

	if len(res) == m != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	for _, cur := range res {
		u, v := cur[0], cur[1]
		if gcd(u, v) != 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

}

func TestSample1(t *testing.T) {
	runSample(t, 5, 6, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 12, false)
}
