package p1455

import "testing"

func runSample(t *testing.T, sent, search string, expect int) {
	res := isPrefixOfWord(sent, search)

	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", sent, search, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "i love eating burger", "burg", 4)
}
