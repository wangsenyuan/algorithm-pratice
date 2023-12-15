package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		land := make([]string, n)
		for i := 0; i < n; i++ {
			land[i] = readString(reader)[:m]
		}
		ok, res := solve(land)
		if !ok {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for _, row := range res {
				buf.WriteString(row)
				buf.WriteByte('\n')
			}
		}
	}
	fmt.Print(buf.String())
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

var dx = []int{-1, -1, 1, 1}
var dy = []int{-1, 1, 1, -1}
var dd = []int{-1, 0, 1, 0, -1}

func solve(land []string) (bool, []string) {
	n := len(land)
	m := len(land[0])
	// n * m + 2个点，每个点，有4条边连接
	tot := n*m + 2

	dist := make([]int, tot+2)
	prev := make([]int, tot+2)
	for i := 0; i < len(dist); i++ {
		dist[i] = -1
		prev[i] = -1
	}

	que := make([]int, 2*tot)
	front, tail := tot, tot

	add := func(u int, d int) {
		if d == 0 {
			tail--
			que[tail] = u
		} else {
			que[front] = u
			front++
		}
	}
	// left side
	add(tot, 0)

	addIfPossible := func(u int, v int, src int) {
		// only when (u, v) is free of cacti in neighborhood
		for i := 0; i < 4; i++ {
			// 是上下左右，不是斜角
			x, y := u+dd[i], v+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && land[x][y] == '#' {
				// bad
				return
			}
		}
		prev[u*m+v] = src
		dist[u*m+v] = dist[src] + 1
		add(u*m+v, 1)
	}

	for tail < front {
		cur := que[tail]
		tail++
		if cur == tot {
			// left one
			for r := 0; r < n; r++ {
				if land[r][0] == '.' {
					addIfPossible(r, 0, tot)
				} else {
					// already have cacti
					prev[r*m+0] = tot
					dist[r*m+0] = 0
					add(r*m+0, 0)
				}
			}
		} else {
			x, y := cur/m, cur%m
			if y == m-1 {
				// it can go right side, and it is the first one to reach
				prev[tot+1] = cur
				dist[tot+1] = dist[cur]
				break
			}

			for i := 0; i < 4; i++ {
				u, v := x+dx[i], y+dy[i]
				if u >= 0 && u < n && v >= 0 && v < m {
					if dist[u*m+v] < 0 {
						if land[u][v] == '.' {
							addIfPossible(u, v, cur)
						} else {
							// always possible to use an existing cacti
							prev[u*m+v] = cur
							dist[u*m+v] = dist[cur]
							add(u*m+v, 0)
						}
					}
				}
			}
		}
	}

	if dist[tot+1] < 0 {
		// no solution
		return false, nil
	}

	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = []byte(land[i])
	}

	it := prev[tot+1]

	for it != tot && it >= 0 {
		u, v := it/m, it%m
		res[u][v] = '#'
		it = prev[it]
	}

	return true, formatAsStringArray(res)
}

func formatAsStringArray(land [][]byte) []string {
	res := make([]string, len(land))
	for i := 0; i < len(land); i++ {
		res[i] = string(land[i])
	}
	return res
}
