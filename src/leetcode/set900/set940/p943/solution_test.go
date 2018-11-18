package p943

import "testing"

func runSample(t *testing.T, A []string, expect string) {
	res := shortestSuperstring(A)
	if len(res) != len(expect) {
		t.Errorf("Sample %v, expect %s, but got %s", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []string{"alex", "loves", "leetcode"}
	expect := "alexlovesleetcode"
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}
	expect := "gctaagttcatgcatc"
	runSample(t, A, expect)
}
