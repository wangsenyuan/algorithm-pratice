package p1966

import "testing"

func runSample(t *testing.T, pats []string, word string, expect int) {
	res := numOfStrings(pats, word)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pats := []string{
		"dojsf", "v", "hetnovaoigv", "ksvqfdojsf", "hetnovaoig", "yskjs", "sfr", "msurztkvppptsluk", "ndxffbkkvejuakduhdcfsdbgbt", "znhdgtwzbnh", "h", "ocaualfiscmbpnfalypmtdppymw", "w", "o", "sfjksvqfdo", "odqvzuc", "bozawheajcmlewptgssueylcxhx", "bno", "jhmarzsphxduvdktzqjiz", "j", "sfrjhetnov", "vxv", "ksvqfd", "ognwvqtadalmav", "yqbspvfwmvhycmghabigl", "qyfaiazgdqaw", "ojsfrj", "ojsfrjhetn", "fdojs", "do", "ovaoig", "k", "vrh", "jsfrjhetnovaoigvgk",
	}
	word := "csfjksvqfdojsfrjhetnovaoigvgk"
	expect := 19
	runSample(t, pats, word, expect)
}
