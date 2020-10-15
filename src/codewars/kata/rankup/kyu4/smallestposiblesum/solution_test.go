package smallestposiblesum

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := Solution(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{6, 9, 21}
	expect := 9
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 21, 55}
	expect := 3
	runSample(t, arr, expect)
}
