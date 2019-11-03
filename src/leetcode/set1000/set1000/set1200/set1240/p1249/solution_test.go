package p1249

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := minRemoveToMakeValid(s)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "(a(b(c)d)", "a(b(c)d)")
}
