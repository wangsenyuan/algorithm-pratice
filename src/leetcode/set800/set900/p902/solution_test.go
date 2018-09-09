package p902

import "testing"

func runSample(t *testing.T, D []string, N int, expect int) {
	res := atMostNGivenDigitSet(D, N)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", D, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	D := []string{"1", "3", "5", "7"}
	N := 100
	expect := 20
	runSample(t, D, N, expect)
}

func TestSample2(t *testing.T) {
	D := []string{"1", "4", "9"}
	N := 1000000000
	expect := 29523
	runSample(t, D, N, expect)
}

func TestPow(t *testing.T) {
	res := pow(2, 3)
	if res != 8 {
		t.Errorf("pow(2, 3) should equal to 8, but got %d", res)
	}
}
