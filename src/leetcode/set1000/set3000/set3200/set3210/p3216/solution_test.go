package p3216

import (
	"testing"
)

func runSample(t *testing.T, horizontal []int, vertical []int, expect int) {
	res := minimumCost(len(horizontal)+1, len(vertical)+1, horizontal, vertical)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	horizontalCut := []int{1, 3}
	verticalCut := []int{5}
	expect := 13
	runSample(t, horizontalCut, verticalCut, expect)
}
