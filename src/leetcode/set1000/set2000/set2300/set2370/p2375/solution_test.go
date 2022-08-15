package p2375

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := smallestNumber(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pattern := "IIIDIDDD"
	expect := "123549876"
	runSample(t, pattern, expect)
}
