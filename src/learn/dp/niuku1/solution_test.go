package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve([]byte(s))

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "", ".")
}

func TestSample2(t *testing.T) {
	runSample(t, "Who love?, Solo..", "Who Love Solo.")
}

func TestSample4(t *testing.T) {
	runSample(t, "66666666664123+Who-32didn't love? Solo32..", "Who Didn T Love Solo.")
}

func TestSample5(t *testing.T) {
	runSample(t, "abc", "Abc.")
}
