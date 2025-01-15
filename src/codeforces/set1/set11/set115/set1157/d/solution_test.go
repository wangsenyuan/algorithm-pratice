package main

import "testing"

func runSample(t *testing.T, n int, k int, expect bool) {
	res := solve(n, k)

	if (len(res) == k) != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	var sum int
	for i := 0; i < k; i++ {
		sum += res[i]
		if i > 0 {
			if res[i] <= res[i-1] || res[i] > 2*res[i-1] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
	if sum != n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 26, 6, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 3, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 9, 4, false)
}

func TestSample5(t *testing.T) {
	runSample(t, 1000000000, 10, true)
}

func TestSample6(t *testing.T) {
	runSample(t, 7, 2, true)
}

func TestSample7(t *testing.T) {
	runSample(t, 1000000000, 1000, true)
}
