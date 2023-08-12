package main

import "testing"

func runSample(t *testing.T, n int, a int, b int, expect bool) {
	res := solve(n, a, b)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var x, y int
	cnt := make(map[int]int)
	for i := 1; i+1 < n; i++ {
		if res[i-1] < res[i] && res[i] > res[i+1] {
			x++
		}
		if res[i-1] > res[i] && res[i] < res[i+1] {
			y++
		}
		cnt[res[i]]++
	}
	cnt[res[0]]++
	cnt[res[n-1]]++

	if a != x || b != y || len(cnt) != n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b := 4, 1, 1
	runSample(t, n, a, b, true)
}

func TestSample2(t *testing.T) {
	n, a, b := 6, 1, 2
	runSample(t, n, a, b, true)
}

func TestSample3(t *testing.T) {
	n, a, b := 6, 4, 0
	runSample(t, n, a, b, false)
}

func TestSample4(t *testing.T) {
	n, a, b := 6, 2, 2
	runSample(t, n, a, b, true)
}

func TestSample5(t *testing.T) {
	n, a, b := 3, 1, 0
	runSample(t, n, a, b, true)
}
