package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, x []int, e int, expect bool) {
	res, can := solve(n, x, e)

	if can != expect {
		t.Errorf("sample %d %v %d expect %t, but got %t", n, x, e, expect, can)
	} else if can {
		var ans float64

		for i := 0; i < n; i++ {
			ans += res[i] * float64(x[i])
		}

		if math.Abs(ans-float64(e)) > 1e-6 {
			t.Errorf("sample %d %v %d, result %v, don't give anser %f, which is wrong with %d", n, x, e, res, ans, e)
		}
	}
}

func TestSample1(t *testing.T) {
	n, e := 1, 2
	x := []int{1}
	runSample(t, n, x, e, false)
}

func TestSample2(t *testing.T) {
	n, e := 3, 3
	x := []int{1, 2, 3}
	runSample(t, n, x, e, true)
}

func TestSample3(t *testing.T) {
	n, e := 4, 4
	x := []int{1, 2, 3, 6}
	runSample(t, n, x, e, true)
}
