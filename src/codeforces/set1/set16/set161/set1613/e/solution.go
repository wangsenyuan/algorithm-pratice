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
		field := make([]string, n)
		for i := 0; i < n; i++ {
			field[i] = readString(reader)[:m]
		}
		res := solve(field)
		for _, r := range res {
			buf.WriteString(r)
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(field []string) []string {
	// 如果某个格子能够控制，且旁边的格子，我们也可以将robot移动到这个格子
	// 那么旁边的格子也能够控制
	// 如果出现了 2 * 2的free格子，那么肯定这几个格子肯定无法控制
	n := len(field)
	m := len(field[0])
	deg := make([][]int, n)

	for i := 0; i < n; i++ {
		deg[i] = make([]int, m)
	}

	lx, ly := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == 'L' {
				lx = i
				ly = j
			} else if field[i][j] == '.' {
				for k := 0; k < 4; k++ {
					u, v := i+dd[k], j+dd[k+1]
					if u >= 0 && u < n && v >= 0 && v < m && field[u][v] != '#' {
						deg[i][j]++
					}
				}
			}
		}
	}
	que := make([]int, n*m)
	var front, end int

	push := func(i int, j int) {
		que[end] = i*m + j
		end++
	}

	pop := func() (int, int) {
		cur := que[front]
		front++
		return cur / m, cur % m
	}

	used := make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, m)
	}

	push(lx, ly)
	used[lx][ly] = true

	for front < end {
		x, y := pop()
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < n && v >= 0 && v < m && field[u][v] != '#' && !used[u][v] {
				deg[u][v]--
				if deg[u][v] <= 1 {
					used[u][v] = true
					push(u, v)
				}
			}
		}
	}

	res := make([]string, n)
	buf := make([]byte, m)
	for i := 0; i < n; i++ {
		copy(buf, field[i])
		for j := 0; j < m; j++ {
			if field[i][j] == '.' && used[i][j] {
				buf[j] = '+'
			}
		}
		res[i] = string(buf)
	}
	return res
}
