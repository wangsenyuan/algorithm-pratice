package main

import "testing"

func runSample(t *testing.T, n, K int, expect bool) {
	res := solve(n, K)

	if expect != (len(res) > 0) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	not := make([]int, 0, n-K)
	var j int
	for i := 1; i <= n; i++ {
		for j < K && res[j] < i {
			j++
		}
		if j < K && res[j] == i {
			continue
		}
		not = append(not, i)
	}

	for _, a := range not {
		for _, b := range res {
			if gcd(a, b) > 1 {
				t.Fatalf("result %v not correct", res)
			}
		}
	}

}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 1, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 1, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 2, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 6, 3, false)
}
