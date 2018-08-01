package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := solve(S)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "elcomew elcome to code jam"
	expect := 1
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "wweellccoommee to code qps jam"
	expect := 256
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "welcome to codejam"
	expect := 0
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := `So you've registered. We sent you a welcoming email, to welcome you to code jam. But it's possible that you still don't feel welcomed to code jam. That's why we decided to name a problem "welcome to code jam." After solving this problem, we hope that you'll feel very welcome. Very welcome, that is, to code jam.`
	expect := 3727
	runSample(t, S, expect)
}
