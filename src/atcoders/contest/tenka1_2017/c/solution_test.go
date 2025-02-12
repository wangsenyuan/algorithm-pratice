package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)
	if len(res) == 0 {
		t.Fatalf("Sample expect a solution, but got nil")
	}
	a, b := 0, 1
	for _, cur := range res {
		// a/b += 1/cur
		na := a*cur + b
		nb := b * cur
		c := gcd(na, nb)
		a = na / c
		b = nb / c
	}
	x := gcd(4, n)
	if a*x != 4 || b*x != n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 3485)
}


func TestSample3(t *testing.T) {
	runSample(t, 4664)
}
