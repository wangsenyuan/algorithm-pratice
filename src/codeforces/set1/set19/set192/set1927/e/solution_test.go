package main

import "testing"

func runSample(t *testing.T, n int, k int) {
	res := solve(n, k)

	var sum int
	var mx int
	var mn int
	for i := 0; i < n; i++ {
		sum += res[i]
		if i >= k {
			sum -= res[i-k]

			mx = max(mx, sum)
			mn = min(mn, sum)
		}
		if i == k-1 {
			mx = sum
			mn = sum
		}
	}
	if mx-mn > 1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 100, 66)
}