package main

import "testing"

func runSample(t *testing.T, B []int, expect bool) {
	b := make([]int, len(B))
	copy(b, B)
	ok, res, splits := solve(b)

	if ok != expect {
		t.Fatalf("sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}

	n := len(B)
	A := make([]int, n*n*10)
	copy(A, B)

	for _, op := range res {
		p, c := op[0], op[1]
		copy(A[p+2:], A[p:])
		A[p] = c
		A[p+1] = c
		n += 2
	}
	var cur int
	for _, d := range splits {
		if !checkTandem(A[cur : cur+d]) {
			t.Fatalf("sample get %v, not correct", A[cur:cur+d])
		}
		cur += d
	}
	if cur != n {
		t.Fatalf("Sample result got wrong length result %v", A[:cur])
	}
}

func checkTandem(arr []int) bool {
	n := len(arr)
	if n%2 != 0 {
		return false
	}
	n /= 2
	for i := 0; i < n; i++ {
		if arr[i] != arr[i+n] {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	A := []int{5, 7}
	expect := false
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 5}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 1, 2, 2, 3}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{3, 2, 1, 1, 2, 3}
	expect := true
	runSample(t, A, expect)
}
