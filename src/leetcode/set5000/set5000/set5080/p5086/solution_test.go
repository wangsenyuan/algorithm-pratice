package p5086

import "testing"

func runSample(t *testing.T, text string, expect string) {
	res := smallestSubsequence(text)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", text, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "cdadabcc", "adbc")
}

func TestSample2(t *testing.T) {
	runSample(t, "abcd", "abcd")
}
func TestSample3(t *testing.T) {
	runSample(t, "ecbacba", "eacb")
}

func TestSample4(t *testing.T) {
	runSample(t, "leetcode", "letcod")
}

func TestSample5(t *testing.T) {
	runSample(t, "cbaacabcaaccaacababa", "abc")
}
