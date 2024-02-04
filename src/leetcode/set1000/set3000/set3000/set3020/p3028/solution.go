package p3028

var dd = []int{-1, 0, 1, 0, -1}

func resultGrid(image [][]int, threshold int) [][]int {
	n := len(image)
	m := len(image[0])

	vis := make([][]int, n)
	sum := make([][]int, n)
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
		sum[i] = make([]int, m)
		cnt[i] = make([]int, m)
		for j := 0; j < m; j++ {
			vis[i][j] = -1
		}
	}

	visit := func(i int, j int) {
		// 是否存在一个以(i, j)为左上定点的区域
		if i+3 > n || j+3 > m {
			return
		}

		var tot int

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if x > 0 && abs(image[i+x][j+y]-image[i+x-1][j+y]) > threshold {
					return
				}
				if y > 0 && abs(image[i+x][j+y]-image[i+x][j+y-1]) > threshold {
					return
				}
				tot += image[i+x][j+y]
			}
		}

		tot /= 9

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				sum[i+x][j+y] += tot
				cnt[i+x][j+y]++
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			visit(i, j)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if cnt[i][j] > 0 {
				sum[i][j] /= cnt[i][j]
			} else {
				sum[i][j] = image[i][j]
			}
		}
	}

	return sum
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
