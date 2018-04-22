package p820

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := minimumLengthEncoding(words)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"time", "me", "bell"}
	runSample(t, words, 10)
}

func TestSample2(t *testing.T) {
	words := []string{"time", "me", "bell", "e"}
	runSample(t, words, 10)
}

func TestSample3(t *testing.T) {
	words := []string{"time", "me", "bell", "ye"}
	runSample(t, words, 13)
}
