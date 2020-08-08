package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		maze := make([]string, n)
		for i := 0; i < n; i++ {
			maze[i], _ = reader.ReadString('\n')
		}
		res := solve(n, m, maze)

		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(n, m int, maze []string) bool {
	good := findTarget(maze, 'G')
	if len(good) == 0 {
		return true
	}

	wall := make([][]bool, n)
	for i := 0; i < n; i++ {
		wall[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			wall[i][j] = maze[i][j] == '#'
		}
	}

	N := n * m
	que := make([]int, N)
	vis := make([]bool, N)
	prison := make([]bool, N)

	block := func(bad []int) bool {
		var front, end int
		que[end] = bad[0]*m + bad[1]
		end++
		vis[bad[0]*m+bad[1]] = true

		for front < end {
			cur := que[front]
			front++

			x, y := cur/m, cur%m

			for k := 0; k < 4; k++ {
				u, v := x+dd[k], y+dd[k+1]
				if u >= 0 && u < n && v >= 0 && v < m && maze[u][v] == 'B' && !vis[u*m+v] {
					que[end] = u*m + v
					end++
					vis[u*m+v] = true
				}
			}
		}

		for i := 0; i < end; i++ {
			cur := que[i]
			x, y := cur/m, cur%m
			for k := 0; k < 4; k++ {
				u, v := x+dd[k], y+dd[k+1]
				if u < 0 || u >= n || v < 0 || v >= m || wall[u][v] || maze[u][v] == 'B' {
					continue
				}
				if maze[u][v] == 'G' {
					// Bad aside with Good, not able to prion bad only
					return false
				}
				wall[u][v] = true
			}
			prison[cur] = true
			vis[cur] = false
		}
		return true
	}

	check := func() bool {

		var front, end int
		que[end] = N - 1
		end++
		vis[N-1] = true
		var cnt int
		for front < end && cnt < len(good) {
			cur := que[front]
			front++
			x, y := cur/m, cur%m

			if wall[x][y] {
				return false
			}

			if maze[x][y] == 'G' {
				cnt++
			}

			for k := 0; k < 4; k++ {
				u, v := x+dd[k], y+dd[k+1]
				if u >= 0 && u < n && v >= 0 && v < m && !vis[u*m+v] && !wall[u][v] {
					que[end] = u*m + v
					end++
					vis[u*m+v] = true
				}
			}
		}
		if cnt < len(good) {
			return false
		}

		for i := 0; i < end; i++ {
			vis[que[i]] = false
		}

		return true
	}

	bad := findTarget(maze, 'B')

	for _, cell := range bad {
		if prison[cell[0]*m+cell[1]] {
			continue
		}
		if !block(cell) {
			return false
		}
	}
	// good person already blocked
	if !check() {
		return false
	}

	return true
}

func findTarget(maze []string, x byte) [][]int {
	var res [][]int

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == x {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
