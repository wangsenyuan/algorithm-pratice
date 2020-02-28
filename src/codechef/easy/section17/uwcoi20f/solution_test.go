package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	n := len(s)

	ask := func(i int) int {
		var j int
		for i < n && s[i] == s[j] {
			i++
			j++
		}
		return j
	}

	res := solve(n, ask)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "heyhello", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "hellohellohello", 3)
}
