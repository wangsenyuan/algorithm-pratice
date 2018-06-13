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
	var g int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &g)

	for i := 0; i < g; i++ {
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		friends := make([][]int, n)

		for j := 0; j < n; j++ {
			friends[j] = make([]int, n)
			scanner.Scan()
			bs := scanner.Bytes()
			pos := 0
			for pos < len(bs) {
				var x int
				pos = readInt(bs, pos, &x) + 1
				friends[j][x-1] = 1
			}
		}
		a, d := solve(n, friends)
		fmt.Printf("%d %f\n", a+1, d)
	}
}

func solve(n int, friends [][]int) (int, float64) {
	f := make([][]int, n)
	INF := 1 << 30

	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if friends[i][j] == 1 {
				f[i][j] = 1
			} else if i != j {
				f[i][j] = INF
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if f[i][j] > f[i][k]+f[k][j] {
					f[i][j] = f[i][k] + f[k][j]
				}
			}
		}
	}
	x, y := -1, 0
	for i := 0; i < n; i++ {
		var tmp int
		for j := 0; j < n; j++ {
			tmp += f[i][j]
		}
		if x == -1 || y > tmp {
			x, y = i, tmp
		}
	}
	return x, float64(y) / float64(n)
}

func solve1(n int, friends [][]int) (int, float64) {

	bfs := func(x int) int {
		ds := make([]int, n)
		for i := 0; i < n; i++ {
			ds[i] = -1
		}
		ds[x] = 0
		que := make([]int, n)
		head, tail := 0, 0
		que[tail] = x
		tail++
		var total int
		for head < tail {
			cur := que[head]
			head++
			for i := 0; i < n; i++ {
				if friends[cur][i] == 1 && ds[i] == -1 {
					ds[i] = ds[cur] + 1
					total += ds[i]
					que[tail] = i
					tail++
				}
			}
		}
		return total
	}

	x, min := -1, 0
	for i := 0; i < n; i++ {
		tmp := bfs(i)
		if x == -1 || tmp < min {
			x, min = i, tmp
		}
	}

	return x, float64(min) / float64(n)
}
