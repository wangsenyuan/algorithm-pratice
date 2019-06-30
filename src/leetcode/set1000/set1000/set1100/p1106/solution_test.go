package p1106

import "testing"

func runSample(t *testing.T, expr string, expect bool) {
	res := parseBoolExpr(expr)
	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", expr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "!(f)", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "|(f,t)", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "&(t,f)", false)
}

func TestSample4(t *testing.T) {
	runSample(t, "|(&(t,f,t),!(t))", false)
}
