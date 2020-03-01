package main

import (
	"fmt"
	"testing"
)

func runSample(t *testing.T, n, m int, rects [][]int, expect uint64) {
	res := solve(n, m, rects)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, rects, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 1
	rects := [][]int{
		{10, 10},
	}
	var expect uint64 = 100
	runSample(t, n, m, rects, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 0
	rects := [][]int{
		{5, 10},
		{5, 5},
	}
	var expect uint64 = 25
	runSample(t, n, m, rects, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 1
	rects := [][]int{
		{1, 1},
		{2, 2},
	}
	var expect uint64 = 4
	runSample(t, n, m, rects, expect)
}

func TestAvlTree(t *testing.T) {
	nums := []int{1, 2, 3, 4, -1, 2, 3, 4, 1, -4, -4, -2, -3, -2, -1}

	var root *Node
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num > 0 {
			root = Insert(root, uint64(num))
		} else {
			root = Delete(root, uint64(-num))
		}

		fmt.Printf("%d => %v\n", num, PreOrder(root))
	}
}
