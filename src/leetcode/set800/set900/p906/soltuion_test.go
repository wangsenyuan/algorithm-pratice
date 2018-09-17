package p906

import "testing"

func TestIsPalindrome(t *testing.T) {
	yes := []uint64{11, 121, 1331, 2442}
	no := []uint64{32, 456, 788}

	for _, num := range yes {
		if !isPalindrome(num) {
			t.Errorf("isPalindrome(%d) should be true", num)
		}
	}

	for _, num := range no {
		if isPalindrome(num) {
			t.Errorf("isPalindrome(%d) should be false", num)
		}
	}
}

func runSample(t *testing.T, L, R string, expect int) {
	res := superpalindromesInRange(L, R)
	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "4", "1000", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "3", "484", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "96915129", "1492347521772", 24)
}

func TestSample4(t *testing.T) {
	runSample(t, "38455498359", "999999999999999999", 45)
}
