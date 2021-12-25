package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)
	for i := 0; i < t; i++ {
		var n, m int
		scanner.Scan()
		bs := scanner.Bytes()
		readInt(bs, readInt(bs, 0, &n)+1, &m)
		A := make([][]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			bs = scanner.Bytes()
			A[j] = make([]int, m)
			for k := 0; k < m; k++ {
				A[j][k] = int(bs[k] - '0')
			}
		}
		res := solve(n, m, A)
		fmt.Println(res)
	}
}

func solve(n int, m int, A [][]int) int {
	f := make([][][][][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([][][][]int, m)
		for j := 0; j < m; j++ {
			f[i][j] = make([][][]int, 2)
			for x := 0; x < 2; x++ {
				f[i][j][x] = make([][]int, 2)
				for y := 0; y < 2; y++ {
					f[i][j][x][y] = make([]int, 10)
				}
			}
		}
	}

	clean := func() {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				for x := 0; x < 2; x++ {
					for y := 0; y < 2; y++ {
						for d := 0; d < 10; d++ {
							f[i][j][x][y][d] = -1
						}
					}
				}
			}
		}
	}

	process := func(p int) int {
		var cal func(i, j, x, y, d int) int
		cal = func(i, j, x, y, d int) int {
			if d > 9 || i < 0 || j < 0 {
				return 0
			}
			if f[i][j][x][y][d] >= 0 {
				return f[i][j][x][y][d]
			}
			v := (A[i][j] + x*p + y*p) % 10
			inc := 0
			if v == d {
				inc = 1
			}

			f[i][j][x][y][d] = inc +
				max([]int{
					cal(i, j, x, y, d+1),
					cal(i-1, j, 0, y, d),
					cal(i-1, j, 1, y, d),
					cal(i, j-1, x, 0, d),
					cal(i, j-1, x, 1, d),
					cal(i-1, j-1, 0, 0, d),
					cal(i-1, j-1, 0, 1, d),
					cal(i-1, j-1, 1, 0, d),
					cal(i-1, j-1, 1, 1, d),
				})
			return f[i][j][x][y][d]
		}
		return max([]int{
			cal(n-1, m-1, 0, 0, 0),
			cal(n-1, m-1, 0, 1, 0),
			cal(n-1, m-1, 1, 0, 0),
			cal(n-1, m-1, 1, 1, 0),
		})
	}

	best := 0
	for p := 0; p < 10; p++ {
		clean()
		best = max([]int{best, process(p)})
	}
	return best
}

func max(xs []int) int {
	res := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > res {
			res = xs[i]
		}
	}
	return res
}
