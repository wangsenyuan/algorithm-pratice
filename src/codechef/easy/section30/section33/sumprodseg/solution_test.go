package main

import "testing"

func runSample(t *testing.T, X, Y int64, expect bool) {
	res := solve(X, Y)

	if expect != (len(res) > 0) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if expect && !check(t, res, X, Y) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func check(t *testing.T, r [][]int64, X, Y int64) bool {
	a := r[0]
	b := r[1]
	if a[0] < b[1] && b[0] < a[1] {
		t.Errorf("result segments %v, %v, intercect", a, b)
		return false
	}
	if !checkSumOrProduct(a, X, Y) {
		t.Errorf("result segment %v not correct, neight sum %d, nor product %d", a, X, Y)
		return false
	}
	if !checkSumOrProduct(b, X, Y) {
		t.Errorf("result segment %v not correct, neight sum %d, nor product %d", a, X, Y)
		return false
	}
	return true
}

func checkSumOrProduct(a []int64, X int64, Y int64) bool {
	if a[0]+a[1] == X {
		return true
	}
	if a[0]*a[1] == Y {
		return true
	}
	return false
}

func TestSample1(t *testing.T) {
	var X, Y int64 = 4, 24
	expect := true
	runSample(t, X, Y, expect)
}

func TestSample2(t *testing.T) {
	var X, Y int64 = 1, 1
	expect := false
	runSample(t, X, Y, expect)
}

func TestSample3(t *testing.T) {
	var X, Y int64 = 3, 30
	expect := true
	runSample(t, X, Y, expect)
}
