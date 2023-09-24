package p2866

import "testing"

func runSample(t *testing.T, maxHeights []int, expect int) {
	res := maximumSumOfHeights(maxHeights)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	maxHeights := []int{5, 3, 4, 1, 1}
	expect := 13
	runSample(t, maxHeights, expect)
}

func TestSample2(t *testing.T) {
	maxHeights := []int{6, 5, 3, 9, 2, 7}
	expect := 22
	runSample(t, maxHeights, expect)
}

func TestSample3(t *testing.T) {
	maxHeights := []int{3, 2, 5, 5, 2, 3}
	expect := 18
	runSample(t, maxHeights, expect)
}
