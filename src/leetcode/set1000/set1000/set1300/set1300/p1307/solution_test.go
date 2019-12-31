package p1307

import "testing"

func runSample(t *testing.T, words []string, result string, expect bool) {
	res := isSolvable(words, result)

	if res != expect {
		t.Errorf("Sample %v %s, expect %t, but got %t", words, result, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"SEND", "MORE"}
	result := "MONEY"
	expect := true
	runSample(t, words, result, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"A", "B"}
	result := "C"
	expect := true
	runSample(t, words, result, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"AA", "B"}
	result := "CD"
	expect := true
	runSample(t, words, result, expect)
}

func TestSample4(t *testing.T) {
	words := []string{"SIX", "SEVEN", "SEVEN"}
	result := "TWENTY"
	expect := true
	runSample(t, words, result, expect)
}

func TestSample5(t *testing.T) {
	words := []string{"THIS", "IS", "TOO"}
	result := "FUNNY"
	expect := true
	runSample(t, words, result, expect)
}

func TestSample6(t *testing.T) {
	words := []string{"LEET", "CODE"}
	result := "POINT"
	expect := false
	runSample(t, words, result, expect)
}
