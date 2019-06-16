package p960

import "testing"

func runSample(t *testing.T, A []string, expect int) {
	res := minDeletionSize(A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []string{"babca", "bbazb"}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []string{"edcba"}
	expect := 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []string{"ghi", "def", "abc"}
	expect := 0
	runSample(t, A, expect)
}
