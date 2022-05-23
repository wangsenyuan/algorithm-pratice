package p2274

import "testing"

func runSample(t *testing.T, bottom, top int, special []int, expect int) {
	res := maxConsecutive(bottom, top, special)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	bottom := 6
	top := 8
	special := []int{7, 6, 8}
	expect := 0
	runSample(t, bottom, top, special, expect)
}

func TestSample2(t *testing.T) {
	bottom := 2
	top := 9
	special := []int{4, 6}
	expect := 3
	runSample(t, bottom, top, special, expect)
}
