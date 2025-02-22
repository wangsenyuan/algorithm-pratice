package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	pipes := make([][]int, m)
	for i := 0; i < m; i++ {
		pipes[i] = readNNums(reader, 3)
	}
	res := solve(n, pipes)
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for _, row := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", row[0], row[1], row[2]))
	}
	fmt.Print(buf.String())
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(n int, pipes [][]int) [][]int {
	next := make([]int, n)
	tank := make([]bool, n)
	D := make([]int, n)
	for i := 0; i < n; i++ {
		tank[i] = true
		next[i] = -1
	}

	for _, pipe := range pipes {
		a, b := pipe[0], pipe[1]
		a--
		b--
		tank[b] = false
		next[a] = b
		D[a] = pipe[2]
	}

	var dfs func(u int, flow int) (int, int)

	dfs = func(u int, flow int) (int, int) {
		if next[u] < 0 {
			// tap
			return u, flow
		}

		return dfs(next[u], min(flow, D[u]))
	}

	var res [][]int

	for i := 0; i < n; i++ {
		if tank[i] {
			j, sz := dfs(i, D[i])
			if sz > 0 {
				res = append(res, []int{i + 1, j + 1, sz})
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
