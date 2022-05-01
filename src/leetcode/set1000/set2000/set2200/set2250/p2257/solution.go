package p2257

import "sort"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	row := make([][]Pair, m)
	col := make([][]Pair, n)

	for i := 0; i < m; i++ {
		row[i] = make([]Pair, 0, 1)
	}

	for j := 0; j < n; j++ {
		col[j] = make([]Pair, 0, 1)
	}

	for _, guard := range guards {
		x, y := guard[0], guard[1]
		row[x] = append(row[x], Pair{y, 0})
		col[y] = append(col[y], Pair{x, 0})
	}

	for _, wall := range walls {
		x, y := wall[0], wall[1]
		row[x] = append(row[x], Pair{y, 1})
		col[y] = append(col[y], Pair{x, 1})
	}

	for i := 0; i < m; i++ {
		sort.Sort(Pairs(row[i]))
	}

	for j := 0; j < n; j++ {
		sort.Sort(Pairs(col[j]))
	}

	grid := make([][]int, m)

	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		cur := row[i]
		var p int

		for j := 0; j < n; j++ {
			if p < len(cur) && cur[p].first == j {
				// a guard or wall at (i, j)
				grid[i][j]++
				p++
				continue
			}
			if p > 0 && cur[p-1].second == 0 || p < len(cur) && cur[p].second == 0 {
				// a guard before
				grid[i][j]++
			}
		}
	}

	for j := 0; j < n; j++ {
		cur := col[j]
		var p int
		for i := 0; i < m; i++ {
			if p < len(cur) && cur[p].first == i {
				grid[i][j]++
				p++
				continue
			}
			if p > 0 && cur[p-1].second == 0 || p < len(cur) && cur[p].second == 0 {
				grid[i][j]++
			}
		}
	}
	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
			}
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].Less(ps[j])
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
