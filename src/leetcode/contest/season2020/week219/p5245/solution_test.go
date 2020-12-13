package p5245

import "testing"

func runSample(t *testing.T, cuboids [][]int, expect int) {
	res := maxHeight(cuboids)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", cuboids, expect, res)
	}
}

func TestSample1(t *testing.T) {
	cubs := [][]int{
		{50, 45, 20}, {95, 37, 53}, {45, 23, 12},
	}
	expect := 190
	runSample(t, cubs, expect)
}

func TestSample2(t *testing.T) {
	cubs := [][]int{
		{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7},
	}
	expect := 102
	runSample(t, cubs, expect)
}
