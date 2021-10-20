package lcp46

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, finalCnt []int, totalNum int64, edges [][]int, plans [][]int, expect []int) {
	res := volunteerDeployment(finalCnt, totalNum, edges, plans)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cnt := []int{1, 16}
	var totalNum int64 = 21
	edges := [][]int{{0, 1}, {1, 2}}
	plans := [][]int{
		{2, 1}, {1, 0}, {3, 0},
	}
	expect := []int{5, 7, 9}
	runSample(t, cnt, totalNum, edges, plans, expect)
}
