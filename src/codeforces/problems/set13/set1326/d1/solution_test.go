package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "a", "a")
}

func TestSample2(t *testing.T) {
	runSample(t, "abcdfdcecba", "abcdfdcba")
}

func TestSample3(t *testing.T) {
	runSample(t, "abbaxyzyx", "xyzyx")
}

func TestSample4(t *testing.T) {
	runSample(t, "codeforces", "c")
}

func TestSample5(t *testing.T) {
	runSample(t, "acbba", "abba")
}

func TestSample6(t *testing.T) {
	runSample(t, "acbbca", "acbbca")
}

func TestSample7(t *testing.T) {
	runSample(t, "aaa", "aaa")
}

func TestSample8(t *testing.T) {
	runSample(t, "abba", "abba")
}

func TestSample9(t *testing.T) {
	runSample(t, "acbykskepnivvinpeks", "skepnivvinpeks")
}
