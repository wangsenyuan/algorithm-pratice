package p844

import "testing"

func runSample(t *testing.T, S, T string, expect bool) {
	res := backspaceCompare(S, T)
	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", S, T, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ab#c", "ad#c", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "ab##", "c#d#", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "a##c", "#a#c", true)
}

func TestSample4(t *testing.T) {
	runSample(t, "a#c", "b", false)
}
