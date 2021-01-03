package p1709

import "testing"

func runSample(t *testing.T, box [][]int, truckSize int, expect int) {
	res := maximumUnits(box, truckSize)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	box := [][]int{{1, 3}, {2, 2}, {3, 1}}
	truck := 4
	expect := 8
	runSample(t, box, truck, expect)
}

func TestSample2(t *testing.T) {
	box := [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	truck := 10
	expect := 91
	runSample(t, box, truck, expect)
}
