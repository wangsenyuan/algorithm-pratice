package p1231

import "testing"

func runSample(t *testing.T, sweetness []int, K int, expect int) {
	res := maximizeSweetness(sweetness, K)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", sweetness, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	sweetness := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	K := 5
	expect := 6
	runSample(t, sweetness, K, expect)
}

func TestSample2(t *testing.T) {
	sweetness := []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	K := 8
	expect := 1
	runSample(t, sweetness, K, expect)
}

func TestSample3(t *testing.T) {
	sweetness := []int{1, 2, 2, 1, 2, 2, 1, 2, 2}
	K := 2
	expect := 5
	runSample(t, sweetness, K, expect)
}
