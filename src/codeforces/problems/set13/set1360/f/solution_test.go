package main

import "testing"

func runSample(t *testing.T, n, m int, A []string, expect bool) {
	a := make([][]byte, n)

	for i := 0; i < n; i++ {
		a[i] = []byte(A[i])
	}

	res := solve(n, m, a)

	if expect != (len(res) == m) {
		t.Fatalf("sample %v, expect %t, but got %s", A, expect, res)
	}

	if !expect {
		return
	}

	for i := 0; i < n; i++ {
		if countDiff(A[i], res) > 1 {
			t.Fatalf("res %s, differ with %s more than 1", res, A[i])
		}
	}
}

func countDiff(a, b string) int {
	var cnt int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt
}

func TestSample1(t *testing.T) {
	n, m := 2, 4
	A := []string{
		"abac",
		"zbab",
	}
	expect := true

	runSample(t, n, m, A, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 4
	A := []string{
		"aaaa",
		"bbbb",
	}
	expect := false

	runSample(t, n, m, A, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	A := []string{
		"baa",
		"aaa",
		"aab",
	}
	expect := true

	runSample(t, n, m, A, expect)
}
