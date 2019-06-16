package main

import "fmt"

func main() {
	root := &TreeNode{5, &TreeNode{2, nil, nil}, &TreeNode{-3, nil, nil}}
	fmt.Println(findFrequentTreeSum(root))
}

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var sumAndCount func(root *TreeNode) int

	freq := make(map[int]int)

	sumAndCount = func(root *TreeNode) int {
		left := 0
		if root.Left != nil {
			left = sumAndCount(root.Left)
		}
		right := 0
		if root.Right != nil {
			right = sumAndCount(root.Right)
		}
		sum := root.Val + left + right
		freq[sum]++
		return sum
	}

	sumAndCount(root)

	mxFreq := 0

	for _, f := range freq {
		if f > mxFreq {
			mxFreq = f
		}
	}

	var res []int

	for v, f := range freq {
		if f == mxFreq {
			res = append(res, v)
		}
	}
	return res
}
