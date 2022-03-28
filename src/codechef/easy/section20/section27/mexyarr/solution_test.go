package main

import "testing"

func runSample(t *testing.T, B []int, expect bool) {
	res := solve(B)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if len(res) > 0 && !check(B, res) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func check(B []int, A []int) bool {
	cnt := make(map[int]int)
	var mex int
	for i := 0; i < len(B); i++ {
		cnt[A[i]]++
		if B[i] >= 0 {
			for cnt[mex] > 0 {
				mex++
			}
			if mex != B[i] {
				return false
			}
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	B := []int{0, -1, 2, -1, 5, -1, -1}
	expect := true
	runSample(t, B, expect)
}

func TestSample2(t *testing.T) {
	B := []int{2, -1, -1}
	expect := false
	runSample(t, B, expect)
}
