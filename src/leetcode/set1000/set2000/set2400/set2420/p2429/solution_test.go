package p2429

import "testing"

func runSample(t *testing.T, num1 int, num2 int, expect int) {
	res := minimizeXor(num1, num2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 5, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 12, 3)
}