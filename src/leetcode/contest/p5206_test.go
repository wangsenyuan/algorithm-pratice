package contest

import "testing"

func runRemoveDuplicates(t *testing.T, s string, k int, expect string) {
	res := removeDuplicates(s, k)
	if res != expect {
		t.Errorf("Sample %s %d, expect %s, but got %s", s, k, expect, res)
	}
}

func TestRD1(t *testing.T) {
	runRemoveDuplicates(t, "abcd", 2, "abcd")
}

func TestRD2(t *testing.T) {
	runRemoveDuplicates(t, "deeedbbcccbdaa", 3, "aa")
}

func TestRD3(t *testing.T) {
	runRemoveDuplicates(t, "pbbcggttciiippooaais", 2, "ps")
}
