package p2337

import "testing"

func runSample(t *testing.T, start, target string, expect bool) {
	res := canChange(start, target)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	start := "_L__R__R_"
	target := "L______RR"
	expect := true

	runSample(t, start, target, expect)
}

func TestSample2(t *testing.T) {
	start := "R_L_"
	target := "__LR"
	expect := false

	runSample(t, start, target, expect)
}

func TestSample3(t *testing.T) {
	start := "_R"
	target := "R_"
	expect := false

	runSample(t, start, target, expect)
}