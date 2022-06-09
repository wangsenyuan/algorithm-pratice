package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	} else if expect {
		sum := make([]int64, len(res))
		for j, rng := range res {
			for i := rng[0] - 1; i < rng[1]; i++ {
				sum[j] += int64(A[i])
			}
		}
		and := sum[0]
		for i := 1; i < len(res); i++ {
			and &= sum[i]
		}
		if and == 0 {
			t.Errorf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 9, 8, 4}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{16, 15}
	expect := false
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 6, 24, 120}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{8, 6, 0, 16, 24, 0, 8}
	expect := true
	runSample(t, A, expect)
}
