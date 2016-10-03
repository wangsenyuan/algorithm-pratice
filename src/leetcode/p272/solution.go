package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	. "../util"
)

func main() {
	root, target, k := readFile("src/leetcode/p272/input2.txt")
	fmt.Println(closestKValues(root, target, k))
}

func readFile(file string) (*TreeNode, float64, int) {
	fp, e := filepath.Abs(file)
	dat, e := ioutil.ReadFile(fp)
	if e != nil {
		panic(e)
	}
	content := string(dat)
	lines := strings.Split(content, "\n")
	root := ParseAsTree(lines[0])
	target, _ := strconv.ParseFloat(lines[1], 64)
	k, _ := strconv.Atoi(lines[2])

	return root, target, k
}

func closestKValues(root *TreeNode, target float64, k int) []int {

	var closestValue func(root *TreeNode) *TreeNode

	closestValue = func(root *TreeNode) *TreeNode {
		p := diffAbs(float64(root.Val), target)
		if p < 0.000001 {
			return root
		}
		var res *TreeNode
		if target < float64(root.Val) && root.Left != nil {
			res = closestValue(root.Left)
		}

		if target > float64(root.Val) && root.Right != nil {
			res = closestValue(root.Right)
		}

		if res != nil {
			px := diffAbs(float64(res.Val), target)
			if px < p {
				return res
			}
		}
		return root
	}
	var pre []*TreeNode

	var predecessors func(root, target *TreeNode, k int) int
	predecessors = func(root, target *TreeNode, k int) int {
		if root == nil || k <= 0 {
			return 0
		}

		if target.Val <= root.Val {
			return predecessors(root.Left, target, k)
		}

		x := predecessors(root.Right, target, k)
		if x == k {
			return x
		}
		pre = append(pre, root)
		return x + 1 + predecessors(root.Left, target, k-1-x)
	}

	var suc []*TreeNode
	var successors func(root, target *TreeNode, k int) int
	successors = func(root, target *TreeNode, k int) int {
		if root == nil || k <= 0 {
			return 0
		}

		if target.Val >= root.Val {
			return successors(root.Right, target, k)
		}

		x := successors(root.Left, target, k)
		if x == k {
			return x
		}
		suc = append(suc, root)
		return x + 1 + successors(root.Right, target, k-1-x)
	}

	theClosest := closestValue(root)
	predecessors(root, theClosest, k)
	suc = append(suc, theClosest)
	successors(root, theClosest, k)
	i, j := 0, 0

	res := make([]int, 0, k)
	for len(res) < k && (i < len(pre) || j < len(suc)){
		if i == len(pre) {
			res = append(res, suc[j].Val)
			j++
		} else if j == len(suc) {
			res = append(res, pre[i].Val)
			i++
		} else {
			a, b := pre[i], suc[j]
			ad, bd := diffAbs(float64(a.Val), target), diffAbs(float64(b.Val), target)
			if ad < bd {
				res = append(res, pre[i].Val)
				i++
			} else {
				res = append(res, suc[j].Val)
				j++
			}
		}
	}

	return res
}

func diffAbs(a, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}
