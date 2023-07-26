package main

import "testing"

func runSample(t *testing.T, A []int) {
	res := solve(A)
	if len(res) < len(A)/2 {
		t.Fatalf("Sample result %v, deleted too much", res)
	}
	var sum int
	for i := 0; i < len(res); i++ {
		if i&1 == 0 {
			sum += res[i]
		} else {
			sum -= res[i]
		}
	}
	if sum != 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 0, 1, 1, 0, 0}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	A := []int{1, 0, 1, 0, 1, 0}
	runSample(t, A)
}

func TestSample3(t *testing.T) {
	A := []int{0, 1, 0, 1, 0, 1}
	runSample(t, A)
}
