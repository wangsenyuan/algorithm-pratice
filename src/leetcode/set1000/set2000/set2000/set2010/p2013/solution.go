package p2013

type DetectSquares struct {
	grid [][]int
}

func Constructor() DetectSquares {
	grid := make([][]int, 1001)
	for i := 0; i < 1001; i++ {
		grid[i] = make([]int, 1001)
	}
	return DetectSquares{grid}
}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	this.grid[x][y]++
}

func (this *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	var cnt int
	for l := 1; l <= 1000; l++ {
		var a, b, c, d int
		if x-l >= 0 {
			a = this.grid[x-l][y]
		}
		if y+l <= 1000 {
			b = this.grid[x][y+l]
		}
		if x+l <= 1000 {
			c = this.grid[x+l][y]
		}
		if y-l >= 0 {
			d = this.grid[x][y-l]
		}
		if a > 0 && b > 0 {
			cnt += a * b * this.grid[x-l][y+l]
		}
		if b > 0 && c > 0 {
			cnt += b * c * this.grid[x+l][y+l]
		}
		if c > 0 && d > 0 {
			cnt += c * d * this.grid[x+l][y-l]
		}
		if a > 0 && d > 0 {
			cnt += a * d * this.grid[x-l][y-l]
		}
	}

	return cnt
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
