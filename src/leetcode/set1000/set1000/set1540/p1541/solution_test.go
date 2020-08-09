package p1541

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minInsertions(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "()()()()()(", 7)
}
