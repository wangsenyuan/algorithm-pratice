package p777

import "testing"

func runSample(t *testing.T, start, end string, expect bool) {
	res := canTransform(start, end)
	if res != expect {
		t.Errorf("sample %s %s, expect %t, but got %t", start, end, expect, res)
	}
}

func TestSample1(t *testing.T) {
	start, end := "XLLR", "LXLX"
	expect := false
	runSample(t, start, end, expect)
}

func TestSample2(t *testing.T) {
	start, end := "RXXLRXRXL", "XRLXXRRLX"
	expect := true
	runSample(t, start, end, expect)
}

func TestSample3(t *testing.T) {
	start, end := "LR", "RL"
	expect := false
	runSample(t, start, end, expect)
}

func TestSample4(t *testing.T) {
	start, end := "RL", "LR"
	expect := false
	runSample(t, start, end, expect)
}
