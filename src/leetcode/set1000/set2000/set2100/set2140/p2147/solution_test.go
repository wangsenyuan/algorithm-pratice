package p2147

import "testing"

func runSample(t *testing.T, corridor string, expect int) {
	res := numberOfWays(corridor)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	corridor := "SSPPSPS"
	expect := 3
	runSample(t, corridor, expect)
}

func TestSample2(t *testing.T) {
	corridor := "PPSPSP"
	expect := 1
	runSample(t, corridor, expect)
}
