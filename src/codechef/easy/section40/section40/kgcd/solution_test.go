package main

import (
	"testing"
)

func runSample(t *testing.T, k int64, A []int, expect int) {
	res := solve(A, k)

	if res[0] != expect || gcd(A[res[1]-1], A[res[2]-1]) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 4, 4, 1, 3}
	var k int64 = 8
	var expect = 2
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 4, 4, 1, 3}
	var k int64 = 10
	var expect = 4
	runSample(t, k, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 4, 4, 1, 3}
	var k int64 = 9
	var expect = 2
	runSample(t, k, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 4, 4, 1, 3}
	var k int64 = 7
	var expect = 1
	runSample(t, k, A, expect)
}
