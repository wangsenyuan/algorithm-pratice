package p833

import "testing"

func runSample(t *testing.T, S string, indexes []int, sources []string, targets []string, expect string) {
	res := findReplaceString(S, indexes, sources, targets)

	if res != expect {
		t.Errorf("sample %s %v %v %v, expect %s, but got %s", S, indexes, sources, targets, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "abcd"
	indexes := []int{0, 2}
	sources := []string{"a", "cd"}
	targets := []string{"eee", "ffff"}
	expect := "eeebffff"
	runSample(t, S, indexes, sources, targets, expect)
}

func TestSample2(t *testing.T) {
	S := "abcd"
	indexes := []int{0, 2}
	sources := []string{"ab", "ec"}
	targets := []string{"eee", "ffff"}
	expect := "eeecd"
	runSample(t, S, indexes, sources, targets, expect)
}
