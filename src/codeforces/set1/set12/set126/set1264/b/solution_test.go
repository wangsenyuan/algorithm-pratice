package main

import "testing"

func runSample(t *testing.T, cnt []int, expect bool) {
	res := solve(cnt)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cnt := []int{2, 2, 2, 1}
	expect := true
	runSample(t, cnt, expect)
}

func TestSample2(t *testing.T) {
	cnt := []int{2, 2, 2, 3}
	expect := false
	runSample(t, cnt, expect)
}
