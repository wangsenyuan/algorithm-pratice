package p1476

import "testing"

func runSample(t *testing.T, houses []int, k int, expect int) {
	res := minDistance(houses, k)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", houses, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{1, 4, 8, 10, 20}
	k := 3
	expect := 5
	runSample(t, H, k, expect)
}

func TestSample2(t *testing.T) {
	H := []int{2, 3, 5, 12, 18}
	k := 2
	expect := 9
	runSample(t, H, k, expect)
}

func TestSample3(t *testing.T) {
	H := []int{7, 4, 6, 1}
	k := 1
	expect := 8
	runSample(t, H, k, expect)
}

func TestSample4(t *testing.T) {
	H := []int{3, 6, 14, 10}
	k := 4
	expect := 0
	runSample(t, H, k, expect)
}

func TestSample5(t *testing.T) {
	H := []int{8, 14, 20, 23, 4, 25}
	k := 3
	expect := 9
	runSample(t, H, k, expect)
}

func TestSample6(t *testing.T) {
	H := []int{178, 40, 129, 145, 189, 134, 68, 84, 67, 162, 73, 80, 133, 101, 71, 6, 165, 53, 126, 180, 24, 1, 119, 196, 65, 153}
	k := 9
	expect := 109
	runSample(t, H, k, expect)
}
