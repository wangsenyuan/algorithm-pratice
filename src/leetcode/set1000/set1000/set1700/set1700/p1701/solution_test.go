package p1701

import "testing"

func runSample(t *testing.T, S string, expect string) {
	res := maximumBinaryString(S)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "000110", "111011")
}

func TestSample2(t *testing.T) {
	runSample(t, "101010111011001101000110010001100001111", "111111111111111111101111111111111111111")
}

func TestSample3(t *testing.T) {
	runSample(t, "11", "11")
}