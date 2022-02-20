package p2178

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, num int64, expect []int64) {
	res := maximumEvenSplit(num)
	sort.Slice(expect, func(i, j int) bool {
		return expect[i] < expect[j]
	})
	if len(expect) != len(res) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var num int64 = 12
	expect := []int64{2, 4, 6}
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	var num int64 = 28
	expect := []int64{6, 8, 2, 12}
	runSample(t, num, expect)
}
