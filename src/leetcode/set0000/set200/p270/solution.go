package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)
import . "../util"

/**
 * Definition for a binary tree node.
 */

func main() {
	fp, e := filepath.Abs("../p270/input1.txt")
	dat, e := ioutil.ReadFile(fp)
	if e != nil {
		panic(e)
	}
	content := string(dat)
	lines := strings.Split(content, "\n")
	root := ParseAsTree(lines[0])
	target, _ := strconv.ParseFloat(lines[1], 64)
	fmt.Println(closestValue(root, target))
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
