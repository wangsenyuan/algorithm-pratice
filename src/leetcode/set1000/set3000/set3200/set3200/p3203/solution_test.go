package p3203

import "testing"

func runSample(t *testing.T, red int, blue int, expect int) {
	res := maxHeightOfTriangle(red, blue)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 4, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 1, 2)
}