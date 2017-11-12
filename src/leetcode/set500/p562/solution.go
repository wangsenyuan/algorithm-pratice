package main

import "fmt"

func main() {
	M := [][]int{
		[]int{0, 1, 0, 1, 1},
		[]int{1, 1, 0, 0, 1},
		[]int{0, 0, 0, 1, 0},
		[]int{1, 0, 1, 1, 1},
		[]int{1, 0, 0, 0, 1},
	}

	fmt.Println(longestLine(M))
}

func longestLine(M [][]int) int {
	n := len(M)

	if n == 0 {
		return 0
	}

	m := len(M[0])

	G := make([][][]int, n)
	for i := 0; i < n; i++ {
		G[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			// 0 up, 1 left, 2 diagonal, 3 anti-diagonal
			G[i][j] = make([]int, 4)
		}
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if M[i][j] == 1 {
				G[i][j][0] = prev(G, i-1, j, 0) + 1
				G[i][j][1] = prev(G, i, j-1, 1) + 1
				G[i][j][2] = prev(G, i-1, j-1, 2) + 1
				G[i][j][3] = prev(G, i-1, j+1, 3) + 1
				tmp := max(G[i][j])
				if tmp > res {
					res = tmp
				}
			}
		}
	}

	return res
}
func max(G []int) int {
	var res = 0
	for k := 0; k < 4; k++ {
		if G[k] > res {
			res = G[k]
		}
	}
	return res
}

func prev(G [][][]int, i int, j int, k int) int {
	if i >= 0 && j >= 0 && j < len(G[i]) {
		return G[i][j][k]
	}
	return 0
}
