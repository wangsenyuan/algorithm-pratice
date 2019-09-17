package p1190

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := reverseParentheses(s)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "(abcd)", "dcba")
}

func TestSample2(t *testing.T) {
	runSample(t, "(u(love)i)", "iloveu")
}

func TestSample3(t *testing.T) {
	runSample(t, "(ed(et(oc))el)", "leetcode")
}

func TestSample4(t *testing.T) {
	runSample(t, "a(bcdefghijkl(mno)p)q", "apmnolkjihgfedcbq")
}
