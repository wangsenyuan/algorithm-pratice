package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generate(nums []int) []*TreeNode {
	if len(nums) == 0 {
		return []*TreeNode{nil}
	}

	var result []*TreeNode
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		left := generate(nums[:i])
		right := generate(nums[i+1:])

		for j := range left {
			for k := range right {
				result = append(result, &TreeNode{x, left[j], right[k]})
			}
		}
	}

	return result
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	nums := make([]int, n)

	for i := range nums {
		nums[i] = i + 1
	}
	return generate(nums)
}

func main() {
	trees := generateTrees(5)

	for i := range trees {
		output(trees[i])
		fmt.Println()
	}

	fmt.Printf("total %d\n", len(trees))
}

func output(root *TreeNode) {
	if root == nil {
		fmt.Print("_")
		return
	}

	fmt.Printf("%d{", root.Val)
	output(root.Left)
	output(root.Right)
	fmt.Printf("}")
}
