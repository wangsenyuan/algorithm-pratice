package main

import "fmt"

func main() {
	var n, m int

	fmt.Scanf("%d %d", &n, &m)

	dcs := make([][]int, n)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		dcs[i] = []int{x, y}
	}

	var path string

	fmt.Scanf("%s", &path)

	res := solve(n, m, dcs, path)

	for i := 0; i < m; i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, m int, dcs [][]int, path string) []int {
	xmap, ymap := make(map[int]int), make(map[int]int)
	a, b, c, d := 0, 0, 0, 0
	dist := 0
	for _, dc := range dcs {
		if dc[0] < 0 {
			a++
		} else if dc[0] > 0 {
			b++
		}
		if dc[1] < 0 {
			c++
		} else if dc[1] > 0 {
			d++
		}

		xmap[dc[0]]++
		ymap[dc[1]]++

		dist += abs(dc[0]) + abs(dc[1])
	}

	res := make([]int, m)
	x, y := 0, 0
	for i := 0; i < m; i++ {
		if path[i] == 'L' {
			tmp1 := xmap[x]
			tmp2 := xmap[x-1]
			b += tmp1
			dist = dist + b - a
			a -= tmp2
			x--
		} else if path[i] == 'R' {
			tmp1 := xmap[x]
			tmp2 := xmap[x+1]
			a += tmp1
			dist = dist + a - b
			b -= tmp2
			x++
		} else if path[i] == 'D' {
			tmp1 := ymap[y]
			tmp2 := ymap[y-1]
			d += tmp1
			dist = dist + d - c
			c -= tmp2
			y--
		} else {
			tmp1 := ymap[y]
			tmp2 := ymap[y+1]

			c += tmp1
			dist = dist + c - d
			d -= tmp2
			y++
		}

		res[i] = dist
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
