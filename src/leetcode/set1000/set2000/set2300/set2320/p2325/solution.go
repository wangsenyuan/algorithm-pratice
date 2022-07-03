package p2325

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = -1
		}
	}
	sz := m * n
	var p int
	a, b, c, d := 0, 0, m-1, n-1
	for p < sz && head != nil {
		for col := b; col <= d && head != nil; col++ {
			res[a][col] = head.Val
			head = head.Next
			p++
		}
		for row := a + 1; row <= c && head != nil; row++ {
			res[row][d] = head.Val
			head = head.Next
			p++
		}
		for col := d - 1; col >= b && head != nil; col-- {
			res[c][col] = head.Val
			head = head.Next
			p++
		}
		for row := c - 1; row > a && head != nil; row-- {
			res[row][b] = head.Val
			head = head.Next
			p++
		}
		a++
		c--
		b++
		d--
	}

	return res
}
