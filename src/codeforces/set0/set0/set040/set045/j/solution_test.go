package main

import "testing"

func runSample(t *testing.T, n int, m int, expect bool) {
	res := solve(n, m)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 && abs(res[i][j]-res[i-1][j]) == 1 {
				t.Fatalf("Sample result %v, not correct", res)
			}
			if j > 0 && abs(res[i][j]-res[i][j-1]) == 1 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, false)
}
