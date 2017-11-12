package main

import "fmt"

func main() {
	grid := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	fmt.Println(numIslands(grid))
}

var neighbors = [2][2]int{[2]int{0, -1}, [2]int{-1, 0}}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	uf := createUF(m * n)

	for i, row := range grid {
		for j := range row {
			if row[j] == '0' {
				continue
			}

			for k := range neighbors {
				a, b := neighbors[k][0]+i, neighbors[k][1]+j
				if a < 0 || a >= m || b < 0 || b >= n {
					continue
				}
				if grid[a][b] == '0' {
					continue
				}
				uf.union(i*n+j, a*n+b)
			}
		}
	}

	sz := make(map[int]bool)
	for i, row := range grid {
		for j := range row {
			if row[j] == '0' {
				continue
			}

			sz[uf.find(i*n+j)] = true
		}
	}
	return len(sz)
}

type UF struct {
	set []int
}

func createUF(n int) *UF {
	set := make([]int, n)
	for i := range set {
		set[i] = i
	}
	return &UF{set}
}

func (uf *UF) find(x int) int {
	p := uf.set[x]
	if p == x {
		return p
	}
	p = uf.find(p)
	uf.set[x] = p
	return p
}

func (uf *UF) union(x, y int) {
	a, b := uf.find(x), uf.find(y)
	uf.set[a] = b
}
