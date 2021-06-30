package p1910

import "testing"

func runSample(t *testing.T, s string, part string, expect string) {
	res := removeOccurrences(s, part)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "daabcbaabcbc"
	p := "abc"
	expect := "dab"
	runSample(t, s, p, expect)
}

func TestSample2(t *testing.T) {
	s := "axxxxyyyyb"
	p := "xy"
	expect := "ab"
	runSample(t, s, p, expect)
}

func TestSample3(t *testing.T) {
	s := "hhvhvaahvahvhvaavhvaasshvahvaln"
	p := "hva"
	expect := "ssln"
	runSample(t, s, p, expect)
}
