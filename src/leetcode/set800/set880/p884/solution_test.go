package p884

import "testing"

func runSample(t *testing.T, S string, K int, expect string) {
	res := decodeAtIndex(S, K)

	if res != expect {
		t.Errorf("Sample %s %d, expect %s, but got %s", S, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "leet2code3", 10, "o")
}

func TestSample2(t *testing.T) {
	runSample(t, "ha22", 5, "h")
}

func TestSample3(t *testing.T) {
	runSample(t, "a2345678999999999999999", 1, "a")
}

func TestSample4(t *testing.T) {
	runSample(t, "a2345678999999999999999", 10000, "a")
}
