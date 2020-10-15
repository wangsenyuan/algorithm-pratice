package snailsort

func Snail(snaipMap [][]int) []int {
	n := len(snaipMap)
	if n == 0 || len(snaipMap[0]) == 0 {
		return []int{}
	}
	res := make([]int, 0, n*n)

	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, n)
	}

	var travel func(x, y int, d int)

	travel = func(x, y int, d int) {
		if x < 0 || y < 0 || x >= n || y >= n || vis[x][y] {
			return
		}
		vis[x][y] = true
		res = append(res, snaipMap[x][y])

		if d == 1 {
			if y+1 == n || vis[x][y+1] {
				travel(x+1, y, 2)
			} else {
				travel(x, y+1, 1)
			}
		} else if d == 2 {
			if x+1 == n || vis[x+1][y] {
				travel(x, y-1, 3)
			} else {
				travel(x+1, y, 2)
			}
		} else if d == 3 {
			if y-1 < 0 || vis[x][y-1] {
				travel(x-1, y, 4)
			} else {
				travel(x, y-1, 3)
			}
		} else {
			if x-1 < 0 || vis[x-1][y] {
				travel(x, y+1, 1)
			} else {
				travel(x-1, y, 4)
			}
		}

	}

	travel(0, 0, 1)

	return res
}
