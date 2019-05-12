package p1041

import "testing"

func runSample(t *testing.T, str string, expect bool) {
	res := isRobotBounded(str)

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", str, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "GL", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "RRGRRGLLLRLGGLGLLGRLRLGLRLRRGLGGLLRRRLRLRLLGRGLGRRRGRLG", false)
}
