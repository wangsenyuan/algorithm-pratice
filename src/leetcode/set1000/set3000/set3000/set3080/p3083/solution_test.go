package p3083

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := minimizeStringValue(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "???"
	expect := "abc"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxy??"
	expect := "abcdefghijklmnopqrstuvwxyaz"
	runSample(t, s, expect)
}
