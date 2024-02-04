package main

import "testing"

func runSample(t *testing.T, a []int, k int) {
	ask := func(i int) int {
		l := i - 1
		r := l + k - 1

		for l < r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}

		var res int
		for j := i - 1; j <= i-1+k-1; j++ {
			res ^= a[j]
		}

		return res
	}

	ans := solve(len(a), k, ask)

	var expect int

	for i := 0; i < len(a); i++ {
		expect ^= a[i]
	}

	if ans != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 2, 5, 1}
	k := 2
	runSample(t, a, k)
}
