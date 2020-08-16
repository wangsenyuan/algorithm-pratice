package p1552

import "testing"

func runSmaple(t *testing.T, arr []int, m int, expect int) {
	res := maxDistance(arr, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pos := []int{1, 2, 3, 4, 7}
	m := 3
	expect := 3
	runSmaple(t, pos, m, expect)
}

func TestSample2(t *testing.T) {
	pos := []int{5, 4, 3, 2, 1, 1000000000}
	m := 2
	expect := 999999999
	runSmaple(t, pos, m, expect)
}
