package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	for i := 1; i <= n; i++ {
		cur := abs(i - res[i-1])
		if !is_prime(cur) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	sort.Ints(res)

	if res[0] != 1 {
		t.Fatalf("Sample result %v not correct", res)
	}
}

func is_prime(x int) bool {
	return x == 2 || x == 3 || x == 5
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	runSample(t, 4, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 8, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 9, true)
}

func TestSample6(t *testing.T) {
	runSample(t, 10, true)
}

func TestSample12(t *testing.T) {
	runSample(t, 11, true)
}
