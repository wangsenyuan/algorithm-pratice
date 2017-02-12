package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 3, 5, 1}
	//nums := []int{1, 3, 2, 3, 1}
	fmt.Println(reversePairs(nums))
}

func reversePairs(nums []int) int {
	var res int
	var root *Node
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		res += search(root, float32(num)/2.0)
		root = build(root, float32(num))
	}

	return res
}

type Node struct {
	val         float32
	eq, lt      int
	left, right *Node
}

func search(root *Node, val float32) int {

	var find func(root *Node, cnt int) int
	find = func(root *Node, cnt int) int {
		if root == nil {
			return cnt
		}
		if val == root.val {
			return root.lt + cnt;
		}

		if val < root.val {
			return find(root.left, cnt)
		}

		return find(root.right, cnt+root.eq+root.lt)
	}

	return find(root, 0)
}

func build(root *Node, val float32) *Node {
	if root == nil {
		return &Node{val, 1, 0, nil, nil}
	}

	if root.val == val {
		root.eq++
	} else if root.val < val {
		root.right = build(root.right, val)
	} else {
		root.left = build(root.left, val)
		root.lt++
	}
	return root

}
