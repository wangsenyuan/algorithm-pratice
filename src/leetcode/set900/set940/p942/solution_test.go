package p942

import "testing"

func runSample(t *testing.T, S string) {
	res := diStringMatch(S)
	nums := make(map[int]bool)
	n := len(S)
	for i := 0; i < n; i++ {
		nums[res[i]] = true
		if S[i] == 'I' && res[i] > res[i+1] {
			t.Errorf("Sample %s, got %v, wrong at %d", S, res, i)
			return
		} else if S[i] == 'D' && res[i] < res[i+1] {
			t.Errorf("Sample %s, got %v, wrong at %d", S, res, i)
			return
		}
	}
	nums[res[n]] = true
	if len(nums) != n+1 {
		t.Errorf("Sample %s, got %v, wrongs numbers", S, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "IDID")
}

func TestSample2(t *testing.T) {
	runSample(t, "III")
}

func TestSample3(t *testing.T) {
	runSample(t, "DDI")
}
