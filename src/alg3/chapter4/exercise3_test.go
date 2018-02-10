package chapter4

import (
	"reflect"
	"testing"
)

func runBottomUpMinimalChange(t *testing.T, nums []int, expect [][]int) {
	res := BottomUpMinimalChange(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("BottomUpMinimalChange sample %v, expect %v, but got %v", nums, expect, res)
	}
}

func TestBottomeUpMinimalChange1(t *testing.T) {
	nums := []int{1}
	expect := [][]int{{1}}
	runBottomUpMinimalChange(t, nums, expect)
}

func TestBottomeUpMinimalChange2(t *testing.T) {
	nums := []int{1, 2}
	expect := [][]int{{1, 2}, {2, 1}}
	runBottomUpMinimalChange(t, nums, expect)
}

func TestBottomeUpMinimalChange3(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := [][]int{{1, 2, 3}, {1, 3, 2}, {3, 1, 2}, {3, 2, 1}, {2, 3, 1}, {2, 1, 3}}
	runBottomUpMinimalChange(t, nums, expect)
}

func TestBottomeUpMinimalChange4(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 4, 2, 3}, {4, 1, 2, 3}, {4, 1, 3, 2}, {1, 4, 3, 2},
		{1, 3, 4, 2}, {1, 3, 2, 4}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 4, 1, 2}, {4, 3, 1, 2}, {4, 3, 2, 1},
		{3, 4, 2, 1}, {3, 2, 4, 1}, {3, 2, 1, 4}, {2, 3, 1, 4},
		{2, 3, 4, 1}, {2, 4, 3, 1}, {4, 2, 3, 1}, {4, 2, 1, 3}, {2, 4, 1, 3}, {2, 1, 4, 3}, {2, 1, 3, 4}}
	runBottomUpMinimalChange(t, nums, expect)
}

func runJohnsonTrotter(t *testing.T, n int, expect [][]int) {
	res := JohnsonTrotter(n)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("JohnsonTrotter sample %d, expect %v, but got %v", n, expect, res)
	}
}

func TestJohnsonTrotter1(t *testing.T) {
	expect := [][]int{{1}}
	runJohnsonTrotter(t, 1, expect)
}

func TestJohnsonTrotter2(t *testing.T) {
	expect := [][]int{{1, 2}, {2, 1}}
	runJohnsonTrotter(t, 2, expect)
}

func TestJohnsonTrotter3(t *testing.T) {
	expect := [][]int{{1, 2, 3}, {1, 3, 2}, {3, 1, 2}, {3, 2, 1}, {2, 3, 1}, {2, 1, 3}}
	runJohnsonTrotter(t, 3, expect)
}

func TestJohnsonTrotter4(t *testing.T) {
	expect := [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 4, 2, 3}, {4, 1, 2, 3}, {4, 1, 3, 2}, {1, 4, 3, 2},
		{1, 3, 4, 2}, {1, 3, 2, 4}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 4, 1, 2}, {4, 3, 1, 2}, {4, 3, 2, 1},
		{3, 4, 2, 1}, {3, 2, 4, 1}, {3, 2, 1, 4}, {2, 3, 1, 4},
		{2, 3, 4, 1}, {2, 4, 3, 1}, {4, 2, 3, 1}, {4, 2, 1, 3}, {2, 4, 1, 3}, {2, 1, 4, 3}, {2, 1, 3, 4}}
	runJohnsonTrotter(t, 4, expect)
}

func runLexicographicPermute(t *testing.T, n int, expect [][]int) {
	res := LexicographicPermute(n)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("JohnsonTrotter sample %d, expect %v, but got %v", n, expect, res)
	}
}

func TestLexicographicPermute1(t *testing.T) {
	expect := [][]int{{1}}
	runLexicographicPermute(t, 1, expect)
}

func TestLexicographicPermute2(t *testing.T) {
	expect := [][]int{{1, 2}, {2, 1}}
	runLexicographicPermute(t, 2, expect)
}

func TestLexicographicPermute3(t *testing.T) {
	expect := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	runLexicographicPermute(t, 3, expect)
}
