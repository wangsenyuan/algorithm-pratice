package p1261

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	nums map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	nums := make(map[int]bool)
	var dfs func(x int, node *TreeNode)

	dfs = func(x int, node *TreeNode) {
		if node == nil {
			return
		}
		nums[x] = true
		dfs(2*x+1, node.Left)
		dfs(2*x+2, node.Right)
	}

	dfs(0, root)

	return FindElements{nums}
}

func (this *FindElements) Find(target int) bool {
	return this.nums[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
