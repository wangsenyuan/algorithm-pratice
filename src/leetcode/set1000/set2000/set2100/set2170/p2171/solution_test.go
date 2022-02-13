package p2171

import "testing"

func runSample(t *testing.T, beans []int, expect int64) {
	res := minimumRemoval(beans)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	beans := []int{4, 1, 6, 5}
	var expect int64 = 4
	runSample(t, beans, expect)
}

func TestSample2(t *testing.T) {
	beans := []int{2, 10, 3, 2}
	var expect int64 = 7
	runSample(t, beans, expect)
}
