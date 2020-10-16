package main

import "testing"

func runSample(t *testing.T, n int, a int, b int, h []int, expect int) {
	H := make([]int, n)
	copy(H, h)
	res := solve(n, a, b, H)

	if len(res) != expect {
		t.Errorf("Sample expect %d steps but got %v", expect, res)
		return
	}

	for i := 0; i < len(res); i++ {
		j := res[i] - 1
		h[j] -= a
		if j > 0 {
			h[j-1] -= b
		}
		if j < n-1 {
			h[j+1] -= b
		}
	}

	for i := 0; i < n; i++ {
		if h[i] >= 0 {
			t.Errorf("Sample result %v, not correct, leave %d alive", res, i)
			return
		}
	}
}

func TestSample1(t *testing.T) {
	n, a, b := 3, 2, 1
	h := []int{2, 2, 2}
	expect := 3
	runSample(t, n, a, b, h, expect)
}

func TestSample2(t *testing.T) {
	n, a, b := 4, 3, 1
	h := []int{1, 4, 1, 1}
	expect := 4
	runSample(t, n, a, b, h, expect)
}

func TestSample3(t *testing.T) {
	n, a, b := 4, 5, 3
	h := []int{4, 2, 4, 2}
	expect := 3
	runSample(t, n, a, b, h, expect)
}

func TestSample4(t *testing.T) {
	n, a, b := 10, 9, 5
	h := []int{12, 14, 11, 11, 14, 14, 12, 15, 14, 12}
	expect := 10
	runSample(t, n, a, b, h, expect)
}
