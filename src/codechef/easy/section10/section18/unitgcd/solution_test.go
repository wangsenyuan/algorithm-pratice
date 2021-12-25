package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)
	mem := make(map[int]int)
	for i := 0; i < len(res); i++ {
		if !check(res[i]) {
			t.Errorf("Sample %d, fail with %v", n, res[i])
		}
		for j := 0; j < len(res[i]); j++ {
			mem[res[i][j]]++
			if mem[res[i][j]] > 1 {
				t.Errorf("Sample %d, fail with %v", n, res[i])
			}
		}
	}

	if len(mem) != n {
		t.Errorf("Sample %d, fail with %v", n, res)
	}
}

func check(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if gcd(nums[i], nums[j]) != 1 {
				return false
			}
		}
	}
	return true
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	runSample(t, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 100)
}

func TestSample3(t *testing.T) {
	runSample(t, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, 101)
}
