package p948

import "testing"

func runSample(t *testing.T, A []int, P int, expect int) {
	res := bagOfTokensScore(A, P)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", A, P, expect, res)
	}
}

func TestSampe1(t *testing.T) {
	A := []int{100}
	P := 50
	expect := 0
	runSample(t, A, P, expect)
}

func TestSampe2(t *testing.T) {
	A := []int{100, 200}
	P := 150
	expect := 1
	runSample(t, A, P, expect)
}

func TestSampe3(t *testing.T) {
	A := []int{100, 200, 300, 400}
	P := 200
	expect := 2
	runSample(t, A, P, expect)
}

func TestSampe4(t *testing.T) {
	A := []int{71, 55, 82}
	P := 54
	expect := 0
	runSample(t, A, P, expect)
}
