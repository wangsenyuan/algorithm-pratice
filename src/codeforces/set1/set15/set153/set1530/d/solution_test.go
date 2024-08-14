package main

import "testing"

func runSample(t *testing.T, a []int, expect []int) {
	x := make([]int, len(a))
	copy(x, a)
	k, res := solve(x)

	cnt := make([]int, 2)
	for i := 0; i < len(a); i++ {
		if a[i] == expect[i] {
			cnt[0]++
		}
		if a[i] == res[i] {
			cnt[1]++
		}
		if res[i] == i+1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	if cnt[0] != cnt[1] || k != cnt[0] {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 2}
	expect := []int{3, 1, 2}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 4, 6, 2, 4, 5, 6}
	expect := []int{6, 4, 7, 2, 3, 5, 1}
	runSample(t, a, expect)
}

// 4 4 5 6 1 4
func TestSample3(t *testing.T) {
	a := []int{4, 4, 5, 6, 1, 4}
	expect := []int{2, 4, 5, 6, 1, 3}
	runSample(t, a, expect)
}
