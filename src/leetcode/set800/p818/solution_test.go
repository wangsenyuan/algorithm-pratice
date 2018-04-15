package p818

import "testing"

func runSample(t *testing.T, target int, expect int) {
	res := racecar(target)

	if res != expect {
		t.Errorf("sample %d, expect %d, but got %d", target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 5)
}
