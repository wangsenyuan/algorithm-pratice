package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	r, c := readTwoNums(reader)
	x, y := readTwoNums(reader)
	mat := make([]string, n)
	for i := 0; i < n; i++ {
		mat[i] = readString(reader)[:m]
	}

	res := solve(mat, r, c, x, y)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(mat []string, r int, c int, x int, y int) int {
	n := len(mat)
	m := len(mat[0])
	r--
	c--

	vis := make([][]bool, n)
	dist := make([][]Pair, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
		dist[i] = make([]Pair, m)
	}

	que := make([]int, 2*m*n)

	front, tail := m*n, m*n

	push := func(i int, j int, a, b int) {
		vis[i][j] = true
		dist[i][j] = Pair{a, b}
		que[front] = i*m + j
		front++
	}

	push1 := func(i, j int, a, b int) {
		vis[i][j] = true
		dist[i][j] = Pair{a, b}
		tail--
		que[tail] = i*m + j
	}

	pop := func() (u int, v int) {
		w := que[tail]
		tail++
		u, v = w/m, w%m
		return
	}

	push(r, c, 0, 0)
	for tail < front {
		a, b := pop()
		cur := dist[a][b]

		for k := 0; k < 4; k++ {
			u, v := a+dd[k], b+dd[k+1]
			if u >= 0 && u < n && v >= 0 && v < m && mat[u][v] == '.' && !vis[u][v] {
				if v == b {
					// same column, added to the tail
					push1(u, v, cur.first, cur.second)
				} else if v < b {
					// move left
					if cur.first+1 <= x {
						push(u, v, cur.first+1, cur.second)
					}
				} else {
					// move right
					if cur.second+1 <= y {
						push(u, v, cur.first, cur.second+1)
					}
				}
			}
		}
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] {
				res++
			}
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}
