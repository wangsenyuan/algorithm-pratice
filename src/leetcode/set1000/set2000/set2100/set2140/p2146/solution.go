package p2146

import "sort"

var dd = []int{-1, 0, 1, 0, -1}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
		}
	}

	que := make([]int, m*n)
	var items []Item

	addItem := func(x, y int) {
		if grid[x][y] >= pricing[0] && grid[x][y] <= pricing[1] {
			items = append(items, Item{dist[x][y], grid[x][y], x, y})
		}
	}

	var front, end int
	que[end] = start[0]*n + start[1]
	end++

	dist[start[0]][start[1]] = 0
	addItem(start[0], start[1])

	for front < end {
		x, y := que[front]/n, que[front]%n
		front++

		for i := 0; i < 4; i++ {
			u, v := x+dd[i], y+dd[i+1]
			if u >= 0 && u < m && v >= 0 && v < n && grid[u][v] > 0 && dist[u][v] < 0 {
				dist[u][v] = dist[x][y] + 1
				addItem(u, v)
				que[end] = u*n + v
				end++
			}
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Less(items[j])
	})
	res := make([][]int, min(len(items), k))

	for i := 0; i < len(items) && i < k; i++ {
		res[i] = []int{items[i].row, items[i].col}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Item struct {
	dist  int
	price int
	row   int
	col   int
}

func (this Item) Less(that Item) bool {
	if this.dist < that.dist {
		return true
	}
	if this.dist == that.dist && this.price < that.price {
		return true
	}
	if this.dist == that.dist && this.price == that.price && this.row < that.row {
		return true
	}
	if this.dist == that.dist && this.price == that.price && this.row == that.row {
		return this.col < that.col
	}
	return false
}
