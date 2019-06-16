package p740

import "testing"

func TestSamples(t *testing.T) {
	samples := []struct {
		nums   []int
		expect int
	}{
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
		{[]int{3, 3, 3, 4, 2}, 9},
	}

	for _, sample := range samples {
		res := deleteAndEarn(sample.nums)
		if res != sample.expect {
			t.Errorf("work on sample %v should get %d, but got %d", sample.nums, sample.expect, res)
		}
	}
}
