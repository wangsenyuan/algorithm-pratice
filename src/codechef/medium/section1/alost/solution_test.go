package main

import "testing"

func runSample(t *testing.T, n int, e int, o int, expect bool) {
	res := solve(n, e, o)

	if expect != (len(res) > 0) {
		t.Fatalf("Sample expect %t, but got %t", expect, len(res) > 0)
	}

	if !expect {
		return
	}
	cnt := make([]int, 2)
	cnt[0]++
	var sum int
	var x, y int
	for i := 0; i < n; i++ {
		sum += res[i]
		sum %= 2
		if sum == 0 {
			x += cnt[0]
			y += cnt[1]
		} else {
			x += cnt[1]
			y += cnt[0]
		}

		cnt[sum]++
	}

	if x != e || o != y {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	e, o := 2, 4
	runSample(t, n, e, o, true)
}

func TestSample2(t *testing.T) {
	n := 3
	e, o := 3, 3
	runSample(t, n, e, o, true)
}

func TestSample3(t *testing.T) {
	n := 3
	e, o := 0, 6
	runSample(t, n, e, o, false)
}
