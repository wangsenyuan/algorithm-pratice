package p863

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, root *TreeNode, target *TreeNode, K int, expect []int) {
	res := distanceK(root, target, K)

	sort.Ints(expect)
	sort.Ints(res)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	str := "[0,1,null,3,2]"
	root := ParseStrAsTree(str)
	target := new(TreeNode)
	target.Val = 2
	K := 1
	expect := []int{1}
	runSample(t, root, target, K, expect)
}

func TestSample2(t *testing.T) {
	str := "[3,5,1,6,2,0,8,null,null,7,4]"
	root := ParseStrAsTree(str)
	target := root.Left
	K := 2
	expect := []int{7, 4, 1}
	runSample(t, root, target, K, expect)
}
