package main

import "testing"

func runSample(t *testing.T, n, x int, P []int, cnt int, expect string) {
	cnt2, res := solve(n, x, P)

	if cnt != cnt2 || res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 9, 3
	P := []int{1, 2, 2, 4, 4, 4, 3, 1}
	cnt := 4
	expect := "aabbbbbba"
	runSample(t, n, x, P, cnt, expect)
}

func TestSample2(t *testing.T) {
	n, x := 1, 1
	P := []int{}
	cnt := 1
	expect := "a"
	runSample(t, n, x, P, cnt, expect)
}
