package main

import "testing"

func runSample(t *testing.T, n int, P []int, expect bool) {
	res := solve(n, P)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if !expect {
		return
	}

	cnt := make([]int, n)
	for i := 0; i < len(res); i++ {
		if P[i] == 0 {
			continue
		}
		x := res[i]
		if x == 'L' {
			cnt[max(0, i-P[i])]++
			cnt[i]--
		} else {
			if i+1 < n {
				cnt[i+1]++
			}
			if i+P[i]+1 < n {
				cnt[i+P[i]+1]--
			}
		}
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			cnt[i] += cnt[i-1]
		}
		if cnt[i] <= 0 {
			t.Fatalf("result %s, not corret, %d not on", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 8
	P := []int{0, 0, 3, 1, 1, 1, 1, 2}
	expect := true
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	P := []int{1, 1}
	expect := true
	runSample(t, n, P, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	P := []int{0, 1}
	expect := false
	runSample(t, n, P, expect)
}
