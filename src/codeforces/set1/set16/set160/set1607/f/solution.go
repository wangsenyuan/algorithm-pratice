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
		readString(reader)
		n, m := readTwoNums(reader)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)[:m]
		}
		res := solve(grid)
		buf.WriteString(fmt.Sprintf("%d %d %d\n", res[0], res[1], res[2]))
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

func solve(grid []string) []int {
	n := len(grid)
	m := len(grid[0])

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	type Pair struct {
		first  int
		second int
	}

	ok := func(r int, c int) bool {
		return r >= 0 && r < n && c >= 0 && c < m
	}

	ans := []int{-1, -1, 0}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if res[r][c] > 0 {
				continue
			}
			nr, nc := r, c
			var dep int
			var path []Pair
			for {
				dep--
				res[nr][nc] = dep
				path = append(path, Pair{nr, nc})

				if grid[nr][nc] == 'L' {
					nc--
				} else if grid[nr][nc] == 'R' {
					nc++
				} else if grid[nr][nc] == 'U' {
					nr--
				} else {
					nr++
				}

				if !ok(nr, nc) || res[nr][nc] != 0 {
					break
				}
			}

			var start = 1

			if ok(nr, nc) {
				if res[nr][nc] < 0 {
					cyc := res[nr][nc] - dep + 1
					for i := 0; i < cyc; i++ {
						p := path[len(path)-1]
						path = path[:len(path)-1]
						res[p.first][p.second] = cyc
					}
				}

				start = res[nr][nc] + 1
			}

			for it := len(path) - 1; it >= 0; it-- {
				cur := path[it]
				res[cur.first][cur.second] = start
				start++
			}

			if res[r][c] > ans[2] {
				ans = []int{r, c, res[r][c]}
			}
		}
	}
	ans[0]++
	ans[1]++
	return ans
}
