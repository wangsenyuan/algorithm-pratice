package main

import "testing"

func runSample(t *testing.T, n int, p int, D []int, expect bool) {
	res := solve(n, p, D)

	if expect == (len(res) == 0) {
		t.Errorf("Sample %d %d %v, expect %t, but got %v", n, p, D, expect, res)
	} else if expect && !checkAnswer(n, p, D, res) {
		t.Errorf("Sample %d %d %v, expect %t, but got %v", n, p, D, expect, res)
	}
}

func checkAnswer(n, p int, D []int, res []int) bool {
	var sum int

	for i := 0; i < n; i++ {
		sum += res[i] * D[i]
	}

	if sum <= p {
		return false
	}

	for i := 0; i < n; i++ {
		if res[i] > 0 {
			if sum-D[i] >= p {
				return false
			}
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	n, p := 2, 10
	D := []int{1, 5}
	expect := false
	runSample(t, n, p, D, expect)
}

func TestSample2(t *testing.T) {
	n, p := 2, 4
	D := []int{1, 5}
	expect := true
	runSample(t, n, p, D, expect)
}

func TestSample3(t *testing.T) {
	n, p := 3, 25
	D := []int{3, 5, 10}
	expect := true
	runSample(t, n, p, D, expect)
}
