package p1439

import "testing"

func runSample(t *testing.T, mat [][]int, k int, expect int) {
	res := kthSmallest(mat, k)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", mat, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 3, 11}, {2, 4, 6},
	}
	k := 5
	expect := 7
	runSample(t, mat, k, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1, 3, 11}, {2, 4, 6},
	}
	k := 9
	expect := 17
	runSample(t, mat, k, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{1, 10, 10}, {1, 4, 5}, {2, 3, 6},
	}
	k := 7
	expect := 9
	runSample(t, mat, k, expect)
}

func TestSample4(t *testing.T) {
	mat := [][]int{
		{1, 1, 10}, {2, 2, 9},
	}
	k := 7
	expect := 12
	runSample(t, mat, k, expect)
}
