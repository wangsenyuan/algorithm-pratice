package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, words []string, expect [][]int) {
	res := palindromePairs(words)

	sort.Sort(Arr(expect))

	sort.Sort(Arr(res))

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample %v, expect %v, but got %v", words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	expect := [][]int{
		{0, 1}, {1, 0}, {3, 2}, {2, 4},
	}
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"bat", "tab", "cat"}
	expect := [][]int{
		{0, 1}, {1, 0},
	}
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"a", ""}
	expect := [][]int{
		{0, 1}, {1, 0},
	}
	runSample(t, words, expect)
}

type Arr [][]int

func (arr Arr) Len() int {
	return len(arr)
}

func (arr Arr) Less(i, j int) bool {
	return arr[i][0] < arr[j][0] || arr[i][0] == arr[j][0] && arr[i][1] < arr[j][1]
}

func (arr Arr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
