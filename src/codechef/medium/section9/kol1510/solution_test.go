package main

import "testing"

func runSample(t *testing.T, K int, A []string, expect string) {
	res := solve(K, A)

	if len(res) != len(expect) {
		t.Errorf("Sample %d %v, expect %s, but got %s", K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := 2
	A := []string{"BG", "GB"}
	expect := "GBG"
	runSample(t, K, A, expect)
}

func TestSample2(t *testing.T) {
	K := 3
	A := []string{"BG", "GBB", "BGB"}
	expect := "BGBB"
	runSample(t, K, A, expect)
}
