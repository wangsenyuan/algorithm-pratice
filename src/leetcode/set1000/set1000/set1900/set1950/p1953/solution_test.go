package p1953

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := minimumPerimeter(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 16)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000, 5040)
}

func runCount(t *testing.T, n int64, expect int64) {
	if count(n) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, count(n))
	}
}

func TestCount1(t *testing.T) {
	runCount(t, 1, 12)
}

func TestCount2(t *testing.T) {
	runCount(t, 2, 60)
}
