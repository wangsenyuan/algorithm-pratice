package p2563

import "testing"

func runSample(t *testing.T, gifts []int, k int, expect int) {
	res := pickGifts(gifts, k)
	if int(res) != expect {
		t.Errorf("Sample expect %d, but gto %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{25, 64, 9, 4, 100}
	k := 4
	expect := 29
	runSample(t, arr, k, expect)
}
