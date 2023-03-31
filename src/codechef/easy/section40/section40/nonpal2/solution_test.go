package main

import "testing"

func runSample(t *testing.T, A string, expect bool) {
	res := solve(A)

	if len(res) == len(A) != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if !expect {
		return
	}

	for i := 1; i < len(A); i++ {
		for j := max(0, i-2); j+1 < i; j++ {
			if res[j] == res[i] {
				t.Fatalf("Sample result %s, not correct", res)
			}
		}
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestSample1(t *testing.T) {
	A := "hello"
	expect := true
	runSample(t, A, expect)
}
