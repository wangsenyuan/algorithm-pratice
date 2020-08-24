package p1562

import "testing"

func runSample(t *testing.T, arr []int, m int, expect int) {
	res := findLatestStep(arr, m)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{3, 5, 1, 2, 4}
	m := 1
	expect := 4
	runSample(t, arr, m, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{3, 1, 5, 4, 2}
	m := 2
	expect := -1
	runSample(t, arr, m, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1}
	m := 1
	expect := 1
	runSample(t, arr, m, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{2, 1}
	m := 2
	expect := 2
	runSample(t, arr, m, expect)
}
