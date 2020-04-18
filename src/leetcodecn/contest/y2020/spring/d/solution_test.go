package d

import "testing"

func runSample(t *testing.T, jump []int, expect int) {
	res := minJump(jump)

	if res != expect {
		t.Errorf("Sample %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	jump := []int{2, 5, 1, 1, 1, 1}
	expect := 3
	runSample(t, jump, expect)
}
