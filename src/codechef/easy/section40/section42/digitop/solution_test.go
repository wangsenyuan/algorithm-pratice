package main

import "testing"

func runSample(t *testing.T, A, B []int, k int64, expect bool) {
	res := solve(A, B, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 9}
	B := []int{9, 1}
	var K int64 = 2
	expect := true
	runSample(t, A, B, K, expect)
}

func TestSample2(t *testing.T) {
	// 顺序也是有关系的
	A := []int{1, 11}
	B := []int{11, 1}
	var K int64 = 2
	expect := false
	runSample(t, A, B, K, expect)
}

func TestSample3(t *testing.T) {
	// 顺序也是有关系的
	A := []int{22, 13, 12, 89}
	B := []int{28, 98, 21, 31}
	var K int64 = 1
	expect := false
	runSample(t, A, B, K, expect)
}
