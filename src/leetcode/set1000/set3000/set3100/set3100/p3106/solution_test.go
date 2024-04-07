package p3106

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := getSmallestString(s, k)

	if res != expect {
		t.Fatalf("Sample %s %d, expect %s, but got %s", s, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "zbbz"
	k := 3
	expect := "aaaz"
	runSample(t, s, k, expect)
}
