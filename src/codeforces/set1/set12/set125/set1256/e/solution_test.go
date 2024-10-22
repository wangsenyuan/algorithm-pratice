package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	tot, cnt, res := solve(a)

	if tot != expect {
		t.Fatalf("Sample expect %d, but got %d, %v", expect, tot, res)
	}

	mn := make([]int, cnt)
	ma := make([]int, cnt)
	for i := range cnt {
		mn[i] = inf
		ma[i] = -inf
	}
	freq := make([]int, cnt)

	for i, x := range a {
		if res[i] <= 0 || res[i] > cnt {
			t.Fatalf("Sample result %v, not valid", res)
		}
		j := res[i] - 1
		mn[j] = min(mn[j], x)
		ma[j] = max(ma[j], x)
		freq[j]++
	}

	var ans int
	for i := range cnt {
		ans += ma[i] - mn[i]
		if freq[i] < 3 {
			t.Fatalf("Sample result %v, not correct, team %d has %d students", res, i+1, freq[i])
		}
	}

	if ans != expect {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 3, 4, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 5, 12, 13, 2, 15}
	expect := 7
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 5, 129, 185, 581, 1041, 1909, 1580, 8150}
	expect := 7486
	runSample(t, a, expect)
}
