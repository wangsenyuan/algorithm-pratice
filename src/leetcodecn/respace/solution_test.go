package respace

import "testing"

func runSample(t *testing.T, dict []string, sent string, expect int) {
	res := respace(dict, sent)

	if res != expect {
		t.Errorf("Sample %v %s, expect %d, but got %d", dict, sent, expect, res)
	}
}

func TestSample1(t *testing.T) {
	dictionary := []string{"looked", "just", "like", "her", "brother"}
	sent := "jesslookedjustliketimherbrother"
	expect := 7
	runSample(t, dictionary, sent, expect)
}
