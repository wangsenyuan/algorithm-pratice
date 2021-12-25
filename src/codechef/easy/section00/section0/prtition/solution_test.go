package main

import "testing"

func runSample(t *testing.T, n int, x int, expect bool) {
	can, res := solve(n, x)
	if can != expect {
		t.Errorf("sample %d %d, expect %t, but got %t", n, x, expect, can)
	} else if can {
		sum0, sum1 := int64(0), int64(0)
		for i := 0; i < n; i++ {
			if res[i] == 0 {
				sum0 += int64(i + 1)
			} else if res[i] == 1 {
				sum1 += int64(i + 1)
			}
		}
		if sum0 != sum1 {
			t.Errorf("sample %d %d, result %v, give wrong patition %d %d", n, x, res, sum0, sum1)
		}
	}
}

func TestSample1(t *testing.T) {
	x, n := 2, 4
	expect := true
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	x, n := 5, 5
	expect := true
	runSample(t, n, x, expect)
}

func TestSample3(t *testing.T) {
	x, n := 1, 2
	expect := false
	runSample(t, n, x, expect)
}
