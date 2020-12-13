package p5612

import "testing"

func runSample(t *testing.T, boxes [][]int, ports int, maxBoxes int, maxWeight int, expect int) {
	res := boxDelivering(boxes, ports, maxBoxes, maxWeight)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	boxes := [][]int{
		{1, 1}, {2, 1}, {1, 1},
	}
	ports := 2
	maxBoxes := 3
	maxWeight := 3
	expect := 4
	runSample(t, boxes, ports, maxBoxes, maxWeight, expect)
}

func TestSample2(t *testing.T) {
	boxes := [][]int{
		{1, 2}, {3, 3}, {3, 1}, {3, 1}, {2, 4},
	}
	ports := 3
	maxBoxes := 3
	maxWeight := 6
	expect := 6
	runSample(t, boxes, ports, maxBoxes, maxWeight, expect)
}

func TestSample3(t *testing.T) {
	boxes := [][]int{
		{1, 4}, {1, 2}, {2, 1}, {2, 1}, {3, 2}, {3, 4},
	}
	ports := 3
	maxBoxes := 6
	maxWeight := 7
	expect := 6
	runSample(t, boxes, ports, maxBoxes, maxWeight, expect)
}

func TestSample4(t *testing.T) {
	boxes := [][]int{
		{2, 4}, {2, 5}, {3, 1}, {3, 2}, {3, 7}, {3, 1}, {4, 4}, {1, 3}, {5, 2},
	}
	ports := 5
	maxBoxes := 5
	maxWeight := 7
	expect := 14
	runSample(t, boxes, ports, maxBoxes, maxWeight, expect)
}

func TestSample5(t *testing.T) {
	boxes := [][]int{
		{1, 1}, {2, 1}, {1, 1}, {2, 1}, {1, 1}, {2, 1},
	}
	ports := 2
	maxBoxes := 10
	maxWeight := 10
	expect := 7
	runSample(t, boxes, ports, maxBoxes, maxWeight, expect)
}
