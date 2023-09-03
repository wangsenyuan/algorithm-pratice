package p2842

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := countKSubsequencesWithMaxBeauty(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "am"
	k := 2
	expect := 1
	runSample(t, s, k, expect)
}
