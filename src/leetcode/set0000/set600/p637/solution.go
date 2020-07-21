package p637

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

	var result []float64

	thisLevel := []*TreeNode{root}

	for len(thisLevel) > 0 {
		var nextLevel []*TreeNode

		sum := 0

		for _, node := range thisLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		result = append(result, float64(sum)/float64(len(thisLevel)))

		thisLevel = nextLevel
	}

	return result
}
