package p2231

import "testing"

func runSample(t *testing.T, expr string, expect string) {
	res := minimizeResult(expr)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "247+38"
	expect := "2(47+38)"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "12+34"
	expect := "1(2+3)4"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "999+999"
	expect := "(999+999)"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "1+1"
	expect := "(1+1)"
	runSample(t, s, expect)
}
