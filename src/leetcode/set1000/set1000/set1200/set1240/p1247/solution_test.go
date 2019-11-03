package p1247

import "testing"

func runSample(t *testing.T, s1, s2 string, expect int) {
	res := minimumSwap(s1, s2)
	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", s1, s2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "xx", "yy", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "xy", "yx", 2)
}

func TestSample3(t *testing.T) {
	runSample(t, "xx", "xy", -1)
}

func TestSample4(t *testing.T) {
	runSample(t, "xxyyxyxyxx", "xyyxyxxxyx", 4)
}
