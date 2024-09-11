package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	res := solve(s)

	if !check(s, res[0], expect[0]) {
		t.Fatalf("Sample result %v, failed at min lis", res)
	}

	if !check(s, res[1], expect[1]) {
		t.Fatalf("Sample result %v, failed at max lis", res)
	}
}

func check(s string, res []int, expect []int) bool {
	n := len(res)
	for i := 0; i < n-1; i++ {
		if s[i] == '<' != (res[i] < res[i+1]) {
			return false
		}
	}

	if getLis(res) != getLis(expect) {
		return false
	}

	sort.Ints(res)

	for i := 0; i < n; i++ {
		if res[i] != i+1 {
			return false
		}
	}
	return true
}

func getLis(arr []int) int {
	n := len(arr)
	res := make([]int, n)
	var p int
	for _, num := range arr {
		j := sort.SearchInts(res[:p], num)
		res[j] = num
		if j == p {
			p++
		}
	}
	return p
}

func TestSample1(t *testing.T) {
	s := "<<"
	expect := [][]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := ">><>><"
	//     6,7,  4,5
	expect := [][]int{
		{5, 4, 3, 7, 2, 1, 6},
		{4, 3, 1, 7, 5, 2, 6},
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := ">>><"
	//     6,7,  4,5
	expect := [][]int{
		{4, 3, 2, 1, 5},
		{5, 4, 2, 1, 3},
	}
	runSample(t, s, expect)
}
