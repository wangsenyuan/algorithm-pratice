package p955

import "testing"

func runSample(t *testing.T, A []string, expect int) {
	res := minDeletionSize(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []string{"ca", "bb", "ac"}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []string{"xc", "yb", "za"}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []string{"zyx", "wvu", "tsr"}
	expect := 3
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []string{"xga", "xfb", "yfa"}
	expect := 1
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []string{"abx", "agz", "bgc", "bfc"}
	expect := 1
	runSample(t, A, expect)
}
