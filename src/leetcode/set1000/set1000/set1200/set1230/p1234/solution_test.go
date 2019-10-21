package p1234

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := balancedString(s)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "WWEQERQWQWWRWWERQWEQ", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "QQWE", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "QQQW", 2)
}

func TestSample4(t *testing.T) {
	runSample(t, "QQQQ", 3)
}
