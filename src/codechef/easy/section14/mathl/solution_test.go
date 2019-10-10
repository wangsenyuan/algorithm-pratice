package main

import "testing"

func bruteForce(n int) int64 {
	res := int64(1)

	for i := 1; i <= n; i++ {
		res *= pow(i, n+1-i)
		res %= MOD
	}

	return res
}

func pow(a int, b int) int64 {
	x := int64(a)
	r := int64(1)

	for b > 0 {
		if b&1 == 1 {
			r *= x
			r %= MOD
		}
		x *= x
		x %= MOD
		b >>= 1
	}

	return r
}

func runSample(t *testing.T, n int) {
	expect := bruteForce(n)

	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 100)
}
