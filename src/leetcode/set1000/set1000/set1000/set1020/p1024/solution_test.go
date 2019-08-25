package p1024

import "testing"

func runSample(t *testing.T, S string, expect string) {
	res := removeOuterParentheses(S)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "(()())(())", "()()()")
}

func TestSample2(t *testing.T) {
	runSample(t, "(()())(())(()(()))", "()()()()(())")
}

func TestSample3(t *testing.T) {
	runSample(t, "()()", "")
}

func TestSample4(t *testing.T) {
	runSample(t, "()", "")
}
