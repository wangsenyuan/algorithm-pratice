package main

import "testing"

func runSample(t *testing.T, n int, A string, expect bool) {
	res := solve(n, A)

	if expect != (res != nil) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if expect {
		if !findMex(res, n) {
			t.Fatalf("Sample result %v, not correct", res)
		}

		for i := 0; i <= n; i++ {
			if findMex(res, i) != (A[i] == '1') {
				t.Fatalf("Sample %v result not correct, expect mex %d (%t)", res, i, A[i] == '1')
			}
		}
	}

}

func findMex(arr []int, mex int) bool {
	for i := 0; i < len(arr); i++ {
		mem := make(map[int]bool)
		var cur int
		for j := i; j < len(arr); j++ {
			mem[arr[j]] = true
			for mem[cur] {
				cur++
			}
			if cur == mex {
				return true
			}
		}
	}
	return false
}

func TestSample1(t *testing.T) {
	n := 2
	A := "111"
	expect := true
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := "110100"
	expect := false
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := "110101"
	expect := true
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	A := "11100111"
	expect := true
	runSample(t, n, A, expect)
}