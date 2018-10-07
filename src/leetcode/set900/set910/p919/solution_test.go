package p919

import "testing"

func runSample(t *testing.T, root *TreeNode, ins []int, expect string) {
	sol := Constructor(root)
	for _, num := range ins {
		sol.Insert(num)
	}
	res := sol.Get_root()

	a := SprintTree(res)

	if a != expect {
		t.Errorf("Sample expect %s, but got %s", expect, a)
	}
}

func TestSample1(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	expect := "1,2,null,null,null"
	runSample(t, root, []int{2}, expect)
}

func TestSample2(t *testing.T) {
	root := ParseAsTree("1,2,3,4,5,6")
	root.Val = 1
	ins := []int{7, 8}
	expect := "1,2,null,null,null"
	runSample(t, root, ins, expect)
}
