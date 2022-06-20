package p2309

import "testing"

func runSample(t *testing.T, num int, k int, expect int) {
	res := minimumNumbers(num, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 58, 9, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 37, 2, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 0, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 3000, 1, 10)
}
