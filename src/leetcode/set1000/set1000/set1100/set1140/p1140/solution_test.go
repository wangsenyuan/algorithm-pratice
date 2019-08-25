package p1140

import "testing"

func runSample(t *testing.T, piles []int, expect int) {
	res := stoneGameII(piles)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", piles, expect, res)
	}
}

func TestSample1(t *testing.T) {
	piles := []int{2, 7, 9, 4, 4}
	expect := 10
	runSample(t, piles, expect)
}

//[8,9,5,4,5,4,1,1,9,3,1,10,5,9,6,2,7,6,6,9]
func TestSample2(t *testing.T) {
	piles := []int{8, 9, 5, 4, 5, 4, 1, 1, 9, 3, 1, 10, 5, 9, 6, 2, 7, 6, 6, 9}
	expect := 56
	runSample(t, piles, expect)
}
