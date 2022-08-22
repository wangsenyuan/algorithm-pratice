package p2385

import "testing"

func runSample(t *testing.T, tree []int, start int, expect int) {
	root := createTree(tree)

	res := amountOfTime(root, start)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", tree, expect, res)
	}
}

func createTree(arr []int) *TreeNode {
	var nodes []*TreeNode

	root := new(TreeNode)
	root.Val = arr[0]
	nodes = append(nodes, root)
	var p int
	for i := 1; i < len(arr); {
		cur := nodes[p]
		p++
		cur.Left = NewNode(arr[i])
		i++
		if i < len(arr) {
			cur.Right = NewNode(arr[i])
			i++
		}
		if cur.Left != nil {
			nodes = append(nodes, cur.Left)
		}
		if cur.Right != nil {
			nodes = append(nodes, cur.Right)
		}
	}

	return root
}

func NewNode(v int) *TreeNode {
	if v < 0 {
		return nil
	}
	node := new(TreeNode)
	node.Val = v
	return node
}

func TestSample1(t *testing.T) {
	arr := []int{1, 5, 3, -1, 4, 10, 6, 9, 2}
	start := 3
	expect := 4
	runSample(t, arr, start, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1}
	start := 1
	expect := 0
	runSample(t, arr, start, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 2, -1, 3, -1, 4, -1, 5}
	start := 2
	expect := 3
	runSample(t, arr, start, expect)
}
