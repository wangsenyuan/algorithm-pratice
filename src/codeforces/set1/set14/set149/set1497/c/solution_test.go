package main

import "testing"

func runSample(t *testing.T, n int, k int) {
	res := solve(n, k)
	if len(res) != k {
		t.Fatalf("Sample result %v, not correct", res)
	}
	x := 1
	var sum int
	for _, num := range res {
		x = lcm(x, num)
		sum += num
	}
	if sum != n || x > n/2 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	n, k := 11, 3
	// 5, 5, 1
	runSample(t, n, k)
}
