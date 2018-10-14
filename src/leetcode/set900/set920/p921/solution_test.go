package p921

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := minAddToMakeValid(S)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "())", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "(((", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "()", 0)
}

func TestSample4(t *testing.T) {
	runSample(t, "()))((", 4)
}

func TestSample5(t *testing.T) {
	runSample(t, "(()))(()", 2)
}
