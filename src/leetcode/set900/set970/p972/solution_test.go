package p972

import "testing"

func runSample(t *testing.T, S, T string, expect bool) {
	res := isRationalEqual(S, T)
	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", S, T, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "0.(52)"
	T := "0.5(25)"
	runSample(t, S, T, true)
}

func TestSample2(t *testing.T) {
	S := "8.123(4567)"
	T := "8.123(4566)"
	runSample(t, S, T, false)
}
