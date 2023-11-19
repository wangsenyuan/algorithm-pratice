package main

import "testing"

func runSample(t *testing.T, n int, k int, x int, expect bool) {
	res := solve(n, k, x)

	if expect != (len(res) == k) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	vis := make(map[int]int)

	for _, cur := range res {
		var y int
		for _, num := range cur {
			vis[num]++
			y ^= num
		}

		if y != x {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	if len(vis) != n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n, k, x := 15, 6, 7
	expect := true
	runSample(t, n, k, x, expect)
}

func TestSample2(t *testing.T) {
	n, k, x := 11, 4, 5
	expect := true
	runSample(t, n, k, x, expect)
}

func TestSample3(t *testing.T) {
	n, k, x := 5, 3, 2
	expect := false
	runSample(t, n, k, x, expect)
}

func TestSample4(t *testing.T) {
	n, k, x := 4, 1, 4
	expect := true
	runSample(t, n, k, x, expect)
}

func TestSample5(t *testing.T) {
	n, k, x := 6, 1, 7
	expect := true
	runSample(t, n, k, x, expect)
}

func TestSample6(t *testing.T) {
	n, k, x := 11, 5, 5
	expect := false
	runSample(t, n, k, x, expect)
}
