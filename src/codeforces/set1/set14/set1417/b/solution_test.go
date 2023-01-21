package main

import (
	"testing"
)

func runSample(t *testing.T, n, T int, A []int, expect int) {
	res := solve(n, T, A)

	c := make([]int, 0, n)
	d := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if res[i] == 1 {
			c = append(c, A[i])
		} else if res[i] == 0 {
			d = append(d, A[i])
		} else {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
	x := f(c, T)
	y := f(d, T)

	if x+y != expect {
		t.Errorf("Sample result %v, get wrong answer %d, aganist %d", res, x+y, expect)
	}
}

func f(arr []int, T int) int {
	cnt := make(map[int]int)

	for _, num := range arr {
		cnt[num]++
	}
	var res int
	for k, c := range cnt {
		if k*2 != T {
			c2 := cnt[T-k]
			res += c * c2
		} else {
			res += c * (c - 1) / 2
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	n, T := 6, 7
	A := []int{1, 2, 3, 4, 5, 6}
	expect := 0
	runSample(t, n, T, A, expect)
}

func TestSample2(t *testing.T) {
	n, T := 3, 6
	A := []int{3, 3, 3}
	expect := 1
	runSample(t, n, T, A, expect)
}
