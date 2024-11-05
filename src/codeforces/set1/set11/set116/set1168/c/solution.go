package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(a, queries)

	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("Shi\n")
		} else {
			buf.WriteString("Fou\n")
		}
	}

	fmt.Print(buf.String())
}

func solve(a []int, queries [][]int) []bool {
	h := bits.Len(uint(slices.Max(a)))
	n := len(a)

	prev := make([][]int, n+1)
	next := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prev[i] = make([]int, h)
		next[i] = make([]int, h)
		for j := 0; j < h; j++ {
			prev[i][j] = -1
			next[i][j] = n
		}
	}
	pos := make([]int, h)
	for j := 0; j < h; j++ {
		pos[j] = -1
	}
	for i, num := range a {
		for d := 0; d < h; d++ {
			if (num>>d)&1 == 1 {
				prev[i][d] = max(prev[i+1][d], pos[d])
				if pos[d] >= 0 {
					for d1 := 0; d1 < h; d1++ {
						prev[i][d1] = max(prev[i][d1], prev[pos[d]][d1])
					}
				}
				pos[d] = i
			}
		}
	}

	for j := 0; j < h; j++ {
		pos[j] = n
	}

	for i := n - 1; i >= 0; i-- {
		num := a[i]
		for d := 0; d < h; d++ {
			if (num>>d)&1 == 1 {
				next[i][d] = min(next[i][d], pos[d])
				if pos[d] < n {
					for d1 := 0; d1 < h; d1++ {
						next[i][d] = min(next[i][d], next[pos[d]][d1])
					}
				}
				pos[d] = i
			}
		}
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		x, y := cur[0], cur[1]
		x--
		y--

		ans[i] = false

		for d := 0; d < h; d++ {
			if next[x][d] == y ||
				prev[y][d] == x ||
				next[x][d] > x && prev[y][d] < y && next[x][d] <= prev[y][d] {
				ans[i] = true
				break
			}
		}
	}

	return ans
}
