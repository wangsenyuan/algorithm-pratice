package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fp, e := filepath.Abs("../p270/input1.txt")
	dat, e := ioutil.ReadFile(fp)
	if e != nil {
		panic(e)
	}
	content := string(dat)
	lines := strings.Split(content, "\n")
	root := parseAsTree(lines[0])
	target, _ := strconv.ParseFloat(lines[1], 64)
	fmt.Println(closestValue(root, target))
}

func parseAsTree(str string) *TreeNode {
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)

	var currLevel []*TreeNode
	var nextLevel []*TreeNode
	nodes := strings.Split(str, ",")

	root := parseNode(nodes[0])
	currLevel = append(currLevel, root)

	pt := 0
	ct := 0
	for i := 1; i < len(nodes); i++ {
		parent := currLevel[pt]
		node := parseNode(nodes[i])
		if node != nil {
			nextLevel = append(nextLevel, node)
		}
		if ct == 0 {
			parent.Left = node
			ct++
		} else {
			parent.Right = node
			pt++
			ct = 0
		}
		if pt == len(currLevel) {
			currLevel = nextLevel
			nextLevel = make([]*TreeNode, 0)
			pt = 0
		}
	}
	return root
}

func parseNode(str string) *TreeNode {
	if str == "null" {
		return nil
	}
	return &TreeNode{parseNum(str), nil, nil}
}
func parseNum(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func closestValue(root *TreeNode, target float64) int {

	var find func(root *TreeNode, prec float64) (int, float64)

	find = func(root *TreeNode, prec float64) (int, float64) {
		val := float64(root.Val)
		p := precision(val, target)
		if p < 0.000001 || (root.Left == nil && root.Right == nil) {
			return root.Val, p
		}

		if target < val {
			if root.Left != nil {
				lr, lp := find(root.Left, min(prec, p))
				if lp < p {
					return lr, lp
				}
			}

			return root.Val, p
		}
		if root.Right != nil {
			rr, rp := find(root.Right, min(prec, p))
			if rp < p {
				return rr, rp
			}
		}
		return root.Val, p
	}

	res, _ := find(root, math.MaxFloat64)

	return res
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func precision(a float64, b float64) float64 {
	if a >= b {
		return a - b
	}
	return b - a
}
