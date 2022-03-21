package p2209

import "testing"

func runSample(t *testing.T, floor string, numCarpets int, carpetLen int, expect int) {
	res := minimumWhiteTiles(floor, numCarpets, carpetLen)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	floor := "10110101"
	numCarpets := 2
	carpetLen := 2
	expect := 2
	runSample(t, floor, numCarpets, carpetLen, expect)
}

func TestSample2(t *testing.T) {
	floor := "11111"
	numCarpets := 2
	carpetLen := 3
	expect := 0
	runSample(t, floor, numCarpets, carpetLen, expect)
}
