package main

import (
	"sort"
	"fmt"
)

var dd = []int{-1, 0, 1, 0, -1}

func cutOffTree(forest [][]int) int {
	n := len(forest)
	m := len(forest[0])
	trees := make([]int, 0, n*m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if forest[i][j] > 1 {
				trees = append(trees, i*m+j)
			}
		}
	}

	sort.Sort(byHeight{forest: forest, n: n, m: m, trees: trees})

	queue := make([]int, n*m)

	findPath := func(x, y, a, b int) int {
		head := 0
		tail := 0
		steps := 0
		visited := make([]bool, n*m)
		queue[tail] = x*m + y
		tail++
		visited[x*m+y] = true

		for tail > head {
			steps++

			oldTail := tail
			for i := head; i < oldTail; i++ {
				c, d := queue[i]/m, queue[i]%m
				for j := 0; j < 4; j++ {
					nc, nd := c+dd[j], d+dd[j+1]
					if nc >= 0 && nc < n && nd >= 0 && nd < m && forest[nc][nd] > 0 && !visited[nc*m+nd] {
						if nc == a && nd == b {
							return steps
						}
						visited[nc*m+nd] = true
						queue[tail] = nc*m + nd
						tail++
					}
				}
			}

			head = oldTail
		}

		return -1
	}

	var ans int
	x, y := 0, 0

	for _, pos := range trees {
		a, b := pos/m, pos%m
		if a != x || b != y {
			tmp := findPath(x, y, a, b)
			if tmp < 0 {
				return -1
			}
			ans += tmp
		}
		x, y = a, b
	}

	return ans
}

type byHeight struct {
	forest [][]int
	n      int
	m      int
	trees  []int
}

func (this byHeight) Len() int {
	return len(this.trees)
}

func (this byHeight) Less(x, y int) bool {
	i, j := this.trees[x], this.trees[y]
	a, b := i/this.m, i%this.m
	c, d := j/this.m, j%this.m
	return this.forest[a][b] < this.forest[c][d]
}

func (this byHeight) Swap(x, y int) {
	this.trees[x], this.trees[y] = this.trees[y], this.trees[x]
}

func main() {
	forest := [][]int{
		[]int{1, 2, 3},
		[]int{0, 0, 4},
		[]int{7, 6, 5},
	}

	/*	forest := [][]int{
			[]int{1, 2, 3},
			[]int{0, 0, 0},
			[]int{7, 6, 5},
		}*/

	fmt.Println(cutOffTree(forest))
}
