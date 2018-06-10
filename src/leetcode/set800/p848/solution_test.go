package p848

import "testing"

func runSample(t *testing.T, S string, shifts []int, expect string) {
	res := shiftingLetters(S, shifts)

	if res != expect {
		t.Errorf("Sample %s %v, expect %s, but got %s", S, shifts, expect, res)
	}
}

func TestSample(t *testing.T) {
	S := "abc"
	shifts := []int{3, 5, 9}
	expect := "rpl"
	runSample(t, S, shifts, expect)
}
