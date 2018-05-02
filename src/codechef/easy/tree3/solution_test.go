package main

import (
	"reflect"
	"sort"
	"testing"
)

type Items [][]int

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	for k := 0; k < 4; k++ {
		if items[i][k] < items[j][k] {
			return true
		}
		if items[i][k] > items[j][k] {
			return false
		}
	}
	return false
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func runSample(t *testing.T, n int, edges [][]int, expect [][]int) {
	res := solve(n, edges)

	for _, arr := range expect {
		sort.Ints(arr[1:])
	}

	for _, arr := range res {
		sort.Ints(arr[1:])
	}

	sort.Sort(Items(expect))
	sort.Sort(Items(res))

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := [][]int{{1, 2, 3, 4}}

	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 5},
		{1, 6},
		{6, 7},
	}

	runSample(t, n, edges, nil)
}
