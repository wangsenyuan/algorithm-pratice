package main

import "testing"

func runSample(t *testing.T, A []int64, expect bool, cnt int) {
	ok, ops := solve(A)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if ok {

		if len(ops) != cnt {
			t.Fatalf("Sample expect %d steps, but got %d steps", cnt, len(ops))
		}
		n := len(A)
		for _, op := range ops {
			l, r := op[0], op[1]
			l--
			r--
			reverse(A[l : r+1])
		}

		if gcd_of_array(A[:n-1]) == 1 || gcd_of_array(A[1:]) == 1 {
			t.Errorf("Sample result %v, not correct", ops)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int64{3, 2, 6}
	expect := true
	cnt := 1
	runSample(t, A, expect, cnt)
}

func TestSample2(t *testing.T) {
	A := []int64{3, 2, 4, 8}
	expect := false
	cnt := 0
	runSample(t, A, expect, cnt)
}

func TestSample3(t *testing.T) {
	A := []int64{6, 15, 9, 18}
	expect := true
	cnt := 0
	runSample(t, A, expect, cnt)
}
