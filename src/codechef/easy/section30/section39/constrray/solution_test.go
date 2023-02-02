package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)
	if n&1 == 1 {
		if len(res) > 0 {
			t.Fatalf("Sample result not correct")
		}
	}

	const INF = 1e9
	var p int
	var s int
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		p += res[i]
		s += res[j]
		if i&1 == 0 {
			if p <= s {
				t.Fatalf("Sample result %v, not correct", res)
			}
		} else {
			if p >= s {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
		if res[i] > INF || res[i] < -INF {
			t.Fatalf("Sample result %v, not correct", res)
		}
		if res[j] > INF || res[j] < -INF {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 100000)
}
