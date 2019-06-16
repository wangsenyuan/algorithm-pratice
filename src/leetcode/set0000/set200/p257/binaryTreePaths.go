package main

import "strconv"

func main() {

}

/**
*Definition for a binary tree node
**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	path := make([]string, 0, 1)

	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)

	if left == nil && right == nil {
		path = updatePath(path, strconv.Itoa(root.Val))
		return path
	}

	for _, p := range left {
		path = updatePath(path, strconv.Itoa(root.Val)+"->"+p)
	}

	for _, p := range right {
		path = updatePath(path, strconv.Itoa(root.Val)+"->"+p)
	}

	return path
}

func updatePath(paths []string, path string) []string {
	if len(paths)+1 >= cap(paths) {
		newPaths := make([]string, len(paths), 2*len(paths))
		copy(newPaths, paths)
		paths = newPaths
	}
	paths = append(paths, path)
	return paths
}
