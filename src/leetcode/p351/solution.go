package p351

//NumberOfPatterns solution
func numberOfPatterns(m int, n int) int {
	if n == 1 {
		return 9
	}
	result = 0

	helper(0, 0, make(map[int]struct{}), m, n, 1)

	helper(0, 1, make(map[int]struct{}), m, n, 1)

	result *= 4

	helper(1, 1, make(map[int]struct{}), m, n, 1)

	return result
}

var directions = [][]int{
	[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1},
	[]int{1, 1}, []int{-1, -1}, []int{-1, 1}, []int{1, -1}, []int{2, 1}, []int{1, 2},
	[]int{-2, 1}, []int{-1, 2}, []int{2, -1}, []int{1, -2}, []int{-2, -1}, []int{-1, -2}}

var result = 0

func helper(i int, j int, visited map[int]struct{}, m int, n int, cnt int) {
	//fmt.Printf("check %d %d\n", i, j)
	if cnt >= m && cnt <= n {
		result++
	}

	if cnt == n {
		return
	}

	visited[i*3+j] = struct{}{}

	for _, d := range directions {
		x, y := i+d[0], j+d[1]
		//fmt.Printf("will check %d %d %b\n", x, y, valid(x, y))
		if !valid(x, y) {
			continue
		}

		if _, ok := visited[x*3+y]; ok {
			//fmt.Println("1")
			x, y = x+d[0], y+d[1]
			if !valid(x, y) {
				continue
			}
			if _, ok2 := visited[x*3+y]; ok2 {
				continue
			}
			helper(x, y, visited, m, n, cnt+1)
		} else {
			//fmt.Println("2")
			helper(x, y, visited, m, n, cnt+1)
		}
	}

	delete(visited, i*3+j)
}

func valid(x, y int) bool {
	return x >= 0 && x < 3 && y >= 0 && y < 3
}
