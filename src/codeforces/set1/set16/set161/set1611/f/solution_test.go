package main

import "testing"

func runSample(t *testing.T, s int, A []int, expect int) {
	res := solve(s, A)
	if expect == 0 != (len(res) == 0) {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	if expect == 0 {
		return
	}
	if res[1]-res[0]+1 != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	var sum = int64(s)
	for i := res[0] - 1; i < res[1]; i++ {
		sum += int64(A[i])
		if sum < 0 {
			t.Fatalf("Sample result %v, not correct at %d", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := 10
	A := []int{-16, 2, -6, 8}
	expect := 3
	runSample(t, s, A, expect)
}

func TestSample2(t *testing.T) {
	s := 1000
	A := []int{-100000, -100000, -100000}
	expect := 0
	runSample(t, s, A, expect)
}

func TestSample3(t *testing.T) {
	s := 0
	A := []int{2, 6, -164, 1, -1, -6543}
	expect := 2
	runSample(t, s, A, expect)
}

func TestSample4(t *testing.T) {
	s := 4
	A := []int{-4, -2}
	expect := 1
	runSample(t, s, A, expect)
}
