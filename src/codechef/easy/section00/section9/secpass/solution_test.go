package main

import "testing"

func runSample(t *testing.T, n int, s string, expect string) {
	res := solve(n, s)
	if res != expect {
		t.Errorf("Sample %d %s, expect %s, but got %s", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, "abc", "abc")
}

func TestSample2(t *testing.T) {
	runSample(t, 5, "thyth", "th")
}

func TestSample3(t *testing.T) {
	runSample(t, 5, "abcbc", "abcbc")
}

func TestSample4(t *testing.T) {
	runSample(t, 5, "aaaaa", "a")
}
