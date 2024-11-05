package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, w := readThreeNums(reader)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = readNNums(reader, int(m))
	}

	res := solve(w, a)

	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

var dd = []int{-1, 0, 1, 0, -1}

const inf = 1 << 60

func solve(w int, a [][]int) int {
	n := len(a)
	m := len(a[0])

	que := make([]int, n*m)

	bfs := func(x int, y int) [][]int {
		dist := make([][]int, n)
		for i := 0; i < n; i++ {
			dist[i] = make([]int, m)
			for j := 0; j < m; j++ {
				dist[i][j] = inf
			}
		}
		var head, tail int
		que[head] = x*m + y
		head++
		dist[x][y] = 0

		for tail < head {
			id := que[tail]
			tail++
			u, v := id/m, id%m
			for i := 0; i < 4; i++ {
				r, c := u+dd[i], v+dd[i+1]
				if r >= 0 && r < n && c >= 0 && c < m && a[r][c] >= 0 && dist[r][c] > dist[u][v]+w {
					dist[r][c] = dist[u][v] + w
					que[head] = r*m + c
					head++
				}
			}
		}

		return dist
	}

	d1 := bfs(0, 0)
	d2 := bfs(n-1, m-1)

	ans := d1[n-1][m-1]

	get := func(dist [][]int) int {
		var res int = inf
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if a[i][j] > 0 && dist[i][j] < inf {
					res = min(res, dist[i][j]+a[i][j])
				}
			}
		}

		return res
	}

	x := get(d1)
	y := get(d2)

	if x < inf && y < inf {
		ans = min(ans, x+y)
	}

	if ans >= inf {
		return -1
	}
	return ans
}
