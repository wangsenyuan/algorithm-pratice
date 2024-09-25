package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := solve(n, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := []int{3, 2}
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{1, 5},
		{5, 6},
	}
	expect := []int{3, 3}
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
		{3, 6},
	}
	// 这里叶子节点是， 4， 5， 6
	// 如果叶子4编号为1，那么可以得到最小值1
	// 要获取到最大值，只要4不编号为1，那么就可以获取到最大值2
	// dp[u] = 和 fp[v]是有关系，但不是简单的获取其中的最小值
	expect := []int{2, 1}
	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 5},
		{3, 6},
		{2, 7},
	}
	expect := []int{2, 1}
	runSample(t, n, edges, expect)
}
