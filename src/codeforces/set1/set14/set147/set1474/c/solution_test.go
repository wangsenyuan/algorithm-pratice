package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	x, res := solve(A)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %d %v", expect, x, res)
	}
	if !expect {
		return
	}
	cnt := make(map[int]int)
	for _, a := range A {
		cnt[a]++
	}

	for _, cur := range res {
		a, b := cur[0], cur[1]
		if a+b != x || cnt[a] == 0 || cnt[b] == 0 {
			t.Fatalf("Sample result %v not correct", res)
		}
		cnt[a]--
		cnt[b]--
		x = max(a, b)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 5, 1, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 8, 8, 64, 64}
	expect := false
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 2, 4}
	expect := false
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 14, 3, 11}
	expect := true
	runSample(t, A, expect)
}
