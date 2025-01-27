package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		cnt[res[i]]++
		if res[i] < 1 || res[i] > n || cnt[res[i]] > 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
		if i > 0 {
			num := res[i] + res[i-1]
			if num%2 == 0 || num == 9 {
				continue
			}
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, true)
}