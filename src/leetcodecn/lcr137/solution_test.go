package lcr137

import "testing"

func runSample(t *testing.T, s string, p string, expect bool) {
	res := articleMatch(s, p)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	article := "aa"
	input := "a"
	expect := false
	runSample(t, article, input, expect)
}

func TestSample2(t *testing.T) {
	article := "aa"
	input := "a*"
	expect := true
	runSample(t, article, input, expect)
}

func TestSample3(t *testing.T) {
	article := "aa"
	input := ".*"
	expect := true
	runSample(t, article, input, expect)
}

func TestSample4(t *testing.T) {
	article := "aab"
	input := "c*a*b"
	expect := true
	runSample(t, article, input, expect)
}

func TestSample5(t *testing.T) {
	article := "mississippi"
	input := "mis*is*p*."
	expect := false
	runSample(t, article, input, expect)
}

func TestSample6(t *testing.T) {
	article := "ab"
	input := ".*"
	expect := true
	runSample(t, article, input, expect)
}

func TestSample7(t *testing.T) {
	article := "aaa"
	input := "aaaa"
	expect := false
	runSample(t, article, input, expect)
}

func TestSample8(t *testing.T) {
	article := "a"
	input := "ab*"
	expect := true
	runSample(t, article, input, expect)
}

func TestSample9(t *testing.T) {
	article := "ab"
	input := ".*c"
	expect := false
	runSample(t, article, input, expect)
}
