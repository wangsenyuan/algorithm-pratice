package p1569

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := numOfWays(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 3, 1}
	expect := 1
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{3, 4, 5, 1, 2}
	expect := 5
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 2, 3}
	expect := 0
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{9, 4, 2, 1, 3, 6, 5, 7, 8, 14, 11, 10, 12, 13, 16, 15, 17, 18}
	expect := 216212978
	runSample(t, arr, expect)
}
