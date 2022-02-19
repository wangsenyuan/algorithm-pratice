package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	bad := make([][]int, k)

	for i := 0; i < k; i++ {
		bad[i] = readNNums(reader, 2)
	}
	q := readNum(reader)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, m, bad, Q)

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, m int, bad [][]int, Q [][]int) []int {
	field := make([][]int, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, m)
	}
	for _, cur := range bad {
		a, b := cur[0], cur[1]
		field[a-1][b-1] = 1
	}
	height := make([][]int, n+1)
	height[n] = make([]int, m)
	for i := n - 1; i >= 0; i-- {
		height[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if field[i][j] == 0 {
				height[i][j] = 1 + height[i+1][j]
			} else {
				height[i][j] = 0
			}
		}
	}

	stack := make([]Pair, m)
	width := make([][]int, n)
	for r := 0; r < n; r++ {
		width[r] = make([]int, n+1)
		var p int
		stack[p] = Pair{height[r][0], 0}
		p++
		for c := 1; c < m; c++ {
			cur := height[r][c]
			last := c
			for p > 0 && stack[p-1].first > cur {
				width[r][r+stack[p-1].first-1] = max(width[r][r+stack[p-1].first-1], c-stack[p-1].second)
				last = stack[p-1].second
				p--
			}
			if p == 0 || stack[p-1].first < cur {
				stack[p] = Pair{height[r][c], last}
				p++
			}
		}
		for p > 0 {
			if r+stack[p-1].first-1 >= 0 {
				width[r][r+stack[p-1].first-1] = max(width[r][r+stack[p-1].first-1], m-stack[p-1].second)
			}
			p--
		}

		for rr := n - 1; rr >= r; rr-- {
			width[r][rr] = max(width[r][rr], width[r][rr+1])
		}
		for rr := r; rr < n; rr++ {
			width[r][rr] *= rr - r + 1
		}
	}

	for l := 2; l <= n; l++ {
		for r := 0; r+l-1 < n; r++ {
			width[r][r+l-1] = max(width[r][r+l-1], width[r][r+l-2])
			width[r][r+l-1] = max(width[r][r+l-1], width[r+1][r+l-1])
		}
	}

	res := make([]int, len(Q))

	for i, cur := range Q {
		a, b := cur[0], cur[1]
		a--
		b--
		res[i] = width[a][b]
	}
	return res
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
