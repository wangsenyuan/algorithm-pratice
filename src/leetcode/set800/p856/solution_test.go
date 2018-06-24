package p856

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := scoreOfParentheses(S)

	if res != expect {
		t.Errorf("Sample %s expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "()", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "(())", 2)
}

func TestSample3(t *testing.T) {
	runSample(t, "()()", 2)
}

func TestSample4(t *testing.T) {
	runSample(t, "(()(()))", 6)
}
