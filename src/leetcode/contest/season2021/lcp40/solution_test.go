package lcp40

import "testing"

func runSample(t *testing.T, arr []int, cnt int, expect int) {
	res := maxmiumScore(arr, cnt)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{9, 7, 1, 4, 9}
	cnt := 1
	expect := 4
	runSample(t, arr, cnt, expect)
}
