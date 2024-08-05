package p3242

type neighborSum struct {
	grid [][]int
	pos  []int
}

func Constructor(grid [][]int) neighborSum {
	n := len(grid)
	pos := make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			pos[grid[i][j]] = i*n + j
		}
	}
	return neighborSum{grid, pos}
}

func (this *neighborSum) AdjacentSum(value int) int {
	n := len(this.grid)
	i, j := this.pos[value]/n, this.pos[value]%n
	var res int
	if i > 0 {
		res += this.grid[i-1][j]
	}
	if i+1 < n {
		res += this.grid[i+1][j]
	}
	if j > 0 {
		res += this.grid[i][j-1]
	}
	if j+1 < n {
		res += this.grid[i][j+1]
	}
	return res
}

func (this *neighborSum) DiagonalSum(value int) int {
	n := len(this.grid)
	i, j := this.pos[value]/n, this.pos[value]%n
	var res int
	if i > 0 && j > 0 {
		res += this.grid[i-1][j-1]
	}
	if i > 0 && j+1 < n {
		res += this.grid[i-1][j+1]
	}
	if i+1 < n && j > 0 {
		res += this.grid[i+1][j-1]
	}
	if i+1 < n && j+1 < n {
		res += this.grid[i+1][j+1]
	}
	return res
}
