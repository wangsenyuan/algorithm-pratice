package p2663

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := smallestBeautifulString(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcz"
	k := 26
	expect := "abda"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "dc"
	k := 4
	expect := ""
	runSample(t, s, k, expect)
}
