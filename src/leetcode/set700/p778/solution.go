package p778

var dd = []int{-1, 0, 1, 0, -1}

func swimInWater(grid [][]int) int {
	n := len(grid)
	que := make([]int, n*n)
	checked := make([]int, n*n)

	check := func(time int) bool {
		if time < grid[0][0] {
			return false
		}
		checked[0] = time
		head, tail := 0, 0
		que[tail] = 0
		tail++
		for head < tail {
			tmp := tail
			for head < tmp {
				v := que[head]
				head++
				if v == n*n-1 {
					return true
				}
				r, c := v/n, v%n
				for j := 0; j < 4; j++ {
					x, y := r+dd[j], c+dd[j+1]
					if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] <= time && checked[x*n+y] != time {
						que[tail] = x*n + y
						tail++
						checked[x*n+y] = time
					}
				}
			}
		}
		return false
	}

	left := 0
	right := n*n + 1
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
