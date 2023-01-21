package main

import "testing"

func runSample(t *testing.T, n int, s string) {
	cnt, res := solve(n, s)
	for i := 1; i <= cnt; i++ {
		var prev = -1
		for j := 0; j < n; j++ {
			if res[j] != i {
				continue
			}
			if prev >= 0 && s[j] == s[prev] {
				t.Fatalf("Sample %s, result %v, not correct", s, res)
			}
			prev = j
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, "11011")
}

func TestSample2(t *testing.T) {
	s := "11011011"
	runSample(t, len(s), s)
}

func TestSample3(t *testing.T) {
	s := "11000"
	runSample(t, len(s), s)
}
