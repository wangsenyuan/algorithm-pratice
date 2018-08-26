package p893

import "testing"

func runSample(t *testing.T, A []string, expect int) {
	res := numSpecialEquivGroups(A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}

}

func TestSample1(t *testing.T) {
	runSample(t, []string{"a", "b", "c", "a", "c", "c"}, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, []string{"aa", "bb", "ab", "ba"}, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, []string{"abc", "acb", "bac", "bca", "cab", "cba"}, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, []string{"abcd", "cdab", "adcb", "cbad"}, 1)
}
