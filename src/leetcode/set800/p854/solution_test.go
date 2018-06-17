package p854

import "testing"

func runSample(t *testing.T, A, B string, expect int) {
	res := kSimilarity(A, B)

	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "ab"
	B := "ba"
	runSample(t, A, B, 1)
}

func TestSample2(t *testing.T) {
	A := "abc"
	B := "bca"
	runSample(t, A, B, 2)
}

func TestSample3(t *testing.T) {
	A := "abac"
	B := "baca"
	runSample(t, A, B, 2)
}

func TestSample4(t *testing.T) {
	A := "aabc"
	B := "abca"
	runSample(t, A, B, 2)
}
