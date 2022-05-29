package p468

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := validIPAddress(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1.0.1."
	expect := "Neither"
	runSample(t, s, expect)
}
