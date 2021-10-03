package p2024

import "testing"

func runSample(t *testing.T, A string, k int, expect int) {
	res := maxConsecutiveAnswers(A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	answerKey := "TTFF"
	k := 2
	expect := 4
	runSample(t, answerKey, k, expect)
}


func TestSample2(t *testing.T) {
	answerKey := "TFFT"
	k := 1
	expect := 3
	runSample(t, answerKey, k, expect)
}



func TestSample3(t *testing.T) {
	answerKey := "TTFTTFTT"
	k := 1
	expect := 5
	runSample(t, answerKey, k, expect)
}
