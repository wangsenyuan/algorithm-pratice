package p843

import "testing"

func runSample(t *testing.T, words []string, master *Master) {
	findSecretWord(words, master)

	if master.cnt > 10 || !master.found {
		t.Errorf("Sample %v, not found", words)
	}
}

func TestSample(t *testing.T) {
	words := []string{"acckzz", "ccbazz", "eiowzz", "abcczz"}
	master := &Master{0, "acckzz", false}
	runSample(t, words, master)
}
