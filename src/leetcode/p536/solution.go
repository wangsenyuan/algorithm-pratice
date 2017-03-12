package main

import "strconv"

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	str2tree("4(2(3)(1))(6(5))")
}

/**
Input: "4(2(3)(1))(6(5))"
Output: return the tree root node representing the following tree:

       4
     /   \
    2     6
   / \   /
  3   1 5
 */

func str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}

	n := len(s)

	i, j, found := findLeftString(s)

	if !found {
		x, _ := strconv.Atoi(s)
		return &TreeNode{x, nil, nil}
	}
	x, _ := strconv.Atoi(s[:i])
	left := str2tree(s[i+1:j])
	var right *TreeNode
	if j < n - 1 {
		right = str2tree(s[j+2:n-1])
	}

	return &TreeNode{x, left, right}
}
func findLeftString(s string) (int, int, bool) {
	var i = 0
	for i < len(s) && s[i] != '(' {
		i++
	}

	if i == len(s) {
		return -1, -1, false
	}

	var j = i
	var level = 0

	for j < len(s) {
		if s[j] == '(' {
			level++
		} else if s[j] == ')' {
			level--
		}
		if level == 0 {
			break
		}
		j++
	}

	if j == len(s) {
		return -1, -1, false
	}

	return i, j, true
}
