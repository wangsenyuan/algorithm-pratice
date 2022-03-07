package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)

	if !check(res) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func check(arr []int) bool {
	n := len(arr)
	cnt := make(map[int]bool)
	for i := 0; i+1 < n; i++ {
		if i%2 == 0 {
			if arr[i+1]%arr[i] != 0 {
				return false
			}
		} else {
			if arr[i+1]%arr[i] == 0 {
				return false
			}
		}
		if arr[i] > 2*n {
			return false
		}
		cnt[arr[i]] = true
	}

	cnt[arr[n-1]] = true

	return arr[n-1] <= 2*n && len(cnt) == n
}

func TestSample1(t *testing.T) {
	runSample(t, 100)
}

func TestSample2(t *testing.T) {
	runSample(t, 101)
}
