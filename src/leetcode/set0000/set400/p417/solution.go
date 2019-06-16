package main

import "fmt"

func main() {
	/*matrix := [][]int{
		[]int{1, 2, 2, 3, 5},
		[]int{3, 2, 3, 4, 4},
		[]int{2, 4, 5, 3, 1},
		[]int{6, 7, 1, 4, 5},
		[]int{5, 1, 1, 2, 4},
	}*/
	matrix := [][]int{
		[]int{1, 1},
		[]int{1, 1},
		[]int{1, 1},
	}

	res := pacificAtlantic(matrix)

	fmt.Println(res)
}

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}

func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	checked := make([][]int, m)

	for i := range checked {
		checked[i] = make([]int, n)
	}

	var flow func(i, j int) int

	flow = func(i, j int) int {
		if checked[i][j] == 3 {
			return checked[i][j]
		}

		res := 0
		if i == 0 || j == 0 {
			res |= 1
		}

		if i == m-1 || j == n-1 {
			res |= 2
		}

		checked[i][j] = -1

		for k := 0; k < len(dx); k++ {
			x, y := i+dx[k], j+dy[k]
			if x < 0 || y < 0 || x == m || y == n {
				continue
			}

			if checked[x][y] < 0 {
				continue
			}

			if matrix[i][j] >= matrix[x][y] {
				res |= flow(x, y)
			}
		}

		checked[i][j] = res

		return res
	}

	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r := flow(i, j)
			if r == 3 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
