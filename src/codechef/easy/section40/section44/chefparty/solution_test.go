package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	sad, res := solve(n)

	if sad != expect {
		t.Fatalf("Sample expect %d, but got %d, %v", expect, sad, res)
	}

	var cnt int

	for _, p := range res {
		g := gcd(p[0], p[1])
		if g == 1 {
			cnt++
		}
	}
	if cnt != expect {
		t.Fatalf("Sample result get %d sad pairs, instead of %d", cnt, expect)
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	n := 2
	expect := 1
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	expect := 1
	runSample(t, n, expect)
}
