package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)
	if res > 0 != expect {
		t.Fatalf("Sample expect %t, but got %d", expect, res)
	}
	if !expect {
		return
	}
	res--
	cur := A[res]

	n := len(A)
	for l, r := res-1, res+1; l >= 0 || r < n; {
		if l >= 0 && A[l] >= cur && r < n && A[r] >= cur {
			t.Fatalf("Sample result %d, stuck at %d %d", res, l, r)
		}
		if l >= 0 && A[l] < cur {
			cur++
			l--
			continue
		}
		if r < n && A[r] < cur {
			cur++
			r++
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 3, 4, 4, 5}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 4, 3, 4, 4}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 5, 4, 3, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{5, 4, 3, 5, 5}
	expect := true
	runSample(t, A, expect)
}
