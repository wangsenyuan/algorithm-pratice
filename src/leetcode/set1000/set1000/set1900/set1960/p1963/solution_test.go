package p1963

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minSwaps(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "]]][[["
	expect := 2
	runSample(t, s, expect)
}
