package p752

import "testing"

func runSample(t *testing.T, deadends []string, target string, expect int) {
	res := openLock(deadends, target)
	if res != expect {
		t.Errorf("open lock(%v, %s) should give %d, but got %d", deadends, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1)
}

func TestSample3(t *testing.T) {
	runSample(t, []string{"0000"}, "8888", -1)
}
