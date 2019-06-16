package p875

import "testing"

func runSample(t *testing.T, piles []int, H int, expect int) {
	res := minEatingSpeed(piles, H)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", piles, H, expect, res)
	}
}

func TestSample1(t *testing.T) {
	piles := []int{3, 6, 7, 11}
	H := 8
	expect := 4
	runSample(t, piles, H, expect)
}

func TestSample2(t *testing.T) {
	piles := []int{30, 11, 23, 4, 20}
	H := 5
	expect := 30
	runSample(t, piles, H, expect)
}

func TestSample3(t *testing.T) {
	piles := []int{30, 11, 23, 4, 20}
	H := 6
	expect := 23
	runSample(t, piles, H, expect)
}

func TestSample4(t *testing.T) {
	piles := []int{312884470}
	H := 968709470
	expect := 1
	runSample(t, piles, H, expect)
}
