package p173

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	arr   []int
	index int
}

func Constructor(root *TreeNode) BSTIterator {
	node := root
	var arr []int

	for node != nil {
		if node.Left == nil {
			arr = append(arr, node.Val)
			node = node.Right
		} else {
			tmp := node.Left
			for tmp.Right != nil && tmp.Right != node {
				tmp = tmp.Right
			}
			if tmp.Right == node {
				arr = append(arr, node.Val)
				node = node.Right
				tmp.Right = nil
			} else {
				tmp.Right = node
				node = node.Left
			}
		}
	}
	return BSTIterator{arr, 0}
}

func (this *BSTIterator) Next() int {
	res := this.arr[this.index]
	this.index++
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.arr)
}
