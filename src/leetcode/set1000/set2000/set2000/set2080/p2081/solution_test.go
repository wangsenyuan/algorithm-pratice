package p2081

import "testing"

func runSample(t *testing.T, k int, n int, expect int64) {
	res := kMirror(k, n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 5, 25)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 7, 499)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 17, 20379000)
}

func TestIsKMirror(t *testing.T) {
	res := isKMirror(4602064, 7)
	if !res {
		t.Errorf("4602064 should be 17-mirror")
	}
}
