package main

import "testing"

func runSample(t *testing.T, n int, l []int, r []int, expect bool) {
	res := solve(n, l, r)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	for i := 0; i < n; i++ {
		var cnt int
		for j := i - 1; j >= 0; j-- {
			if res[j] > res[i] {
				cnt++
			}
		}
		if cnt != l[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		cnt = 0
		for j := i + 1; j < n; j++ {
			if res[j] > res[i] {
				cnt++
			}
		}
		if cnt != r[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	l := []int{0, 0, 1, 1, 2}
	r := []int{2, 0, 1, 0, 0}
	expect := true
	runSample(t, n, l, r, expect)
}
func TestSample2(t *testing.T) {
	n := 4
	l := []int{0, 0, 2, 0}
	r := []int{1, 1, 1, 1}
	expect := false
	runSample(t, n, l, r, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	l := []int{0, 0, 0}
	r := []int{0, 0, 0}
	expect := true
	runSample(t, n, l, r, expect)
}
