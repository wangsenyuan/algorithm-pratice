package p742

import "testing"

func runSample(t *testing.T, tree *TreeNode, k int, expect []int) {
	res := findClosestLeaf(tree, k)
	su := false
	for _, cand := range expect {
		if res == cand {
			su = true
			break
		}
	}
	if !su {
		t.Errorf("the clostest node of %d in tree %s, should be one of %v, but got %d", k, tree.String(), expect, res)
	}
}

func TestSample1(t *testing.T) {
	tree := ParseAsTree("[1,3,2]")
	runSample(t, tree, 1, []int{2, 3})
}

func TestSample2(t *testing.T) {
	tree := ParseAsTree("[1]")
	runSample(t, tree, 1, []int{1})
}

func TestSample3(t *testing.T) {
	tree := ParseAsTree("[1,2,3,4,null,null,null,5,null,6]")
	runSample(t, tree, 2, []int{3})
}

func TestSample4(t *testing.T) {
	tree := ParseAsTree("[1,2,3,4,null,null,7,8,9,null,null,12,null,null,null,null,13,null,14]")
	runSample(t, tree, 8, []int{9})
}

func TestSample5(t *testing.T) {
	tree := ParseAsTree("[1,2]")
	runSample(t, tree, 1, []int{2})
}
