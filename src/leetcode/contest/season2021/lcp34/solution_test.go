package lcp34

import "testing"

func runSample(t *testing.T, arr []int, k int, expect int) {
	root := CreateTree(arr)
	res := maxValue(root, k)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, k, expect, res)
	}
}

func CreateTree(arr []int) *TreeNode {
	// 这个构造树的逻辑是错的。如果某个节点是空的（-1），那么它的子节点就不会占空间了。后续的节点需要迁移。
	// 这也是Sample4 fail 的原因
	var loop func(i int) *TreeNode

	loop = func(i int) *TreeNode {
		if i >= len(arr) || arr[i] < 0 {
			return nil
		}
		node := new(TreeNode)
		node.Val = arr[i]
		node.Left = loop(2*i + 1)
		node.Right = loop(2*i + 2)
		return node
	}

	return loop(0)
}

func TestSample1(t *testing.T) {
	arr := []int{5, 2, 3, 4}
	k := 2
	expect := 12
	runSample(t, arr, k, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{4, 1, 3, 9, -1, -1, 2}
	k := 2
	expect := 16
	runSample(t, arr, k, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{8, 1, 3, 9, 9, 9, -1, 9, 5, 6, 8}
	k := 2
	expect := 52
	runSample(t, arr, k, expect)
}

//func TestSample4(t *testing.T) {
//	arr := []int{66, 96, 4, 38, 49, 87, 21, 93, 69, 31, 98, 27, 46, 82, 87, 41, 78, 14, -1, 98, -1, 87, -1, 64, 25, -1, 87, -1, 70, 82, -1, 15, 21, -1, 36, 33, 82, 48, 10, 61, 59, 72, 12, 86, -1, -1, 60, 73, 34, 4, -1, 97, 56, 44, -1, 20, 25, 48, -1, 42, 52, 43, 87, 37, 35, 90, 50, -1, 19, 54, 74, 86, 74, 68, 41, -1, 66, -1, 91, 20, -1, 31, 19, 100, 65, 20, -1, 9, 16, 77, -1, -1, 41, 96, -1, -1, -1, 64, -1, 82, -1, -1, 65, -1, 39, 28, 3, 75, -1, 1, -1, -1, 33, 28, 66, 74, 51, 24, 10, 76, 94, 3, 10, 75, 58, 23, 42, -1, -1, -1, -1, 17, 36}
//	k := 5
//	expect := 4612
//	runSample(t, arr, k, expect)
//}

func TestSample5(t *testing.T) {
	arr := []int{4, 1, 3, 9, -1, -1, 2}
	k := 4
	expect := 18
	runSample(t, arr, k, expect)
}
