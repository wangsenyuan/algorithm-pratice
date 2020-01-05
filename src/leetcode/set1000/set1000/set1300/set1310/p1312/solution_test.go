package p1312

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minInsertions(s)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "zzazz", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "mbadm", 2)
}

func TestSample3(t *testing.T) {
	runSample(t, "leetcode", 5)
}

func TestSample4(t *testing.T) {
	runSample(t, "no", 1)
}
