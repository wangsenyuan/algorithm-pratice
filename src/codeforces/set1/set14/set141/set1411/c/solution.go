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
		n, k := readTwoNums(reader)

		rocks := make([][]int, k)

		for i := 0; i < k; i++ {
			rocks[i] = readNNums(reader, 2)
		}

		res := solve(n, rocks)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, rocks [][]int) int {
	// (a, b) =>
	m := len(rocks)
	// 如果形成了一个环，那么就是size(cycle) + 1
	// 否则的话就是
	row := make([]int, n)
	col := make([]int, n)
	for i := 0; i < n; i++ {
		row[i] = -1
		col[i] = -1
	}
	var good int
	for i := 0; i < m; i++ {
		rock := rocks[i]
		rock[0]--
		rock[1]--
		if rock[0] == rock[1] {
			good++
		}
		row[rock[0]] = i
		col[rock[1]] = i
	}

	vis := make([]int, m)

	par := make([]int, n)

	var cycles []int

	var dfs func(u int, cur int)

	dfs = func(u int, cur int) {
		vis[u] = cur
		rock := rocks[u]
		r, c := rock[0], rock[1]
		if r == c {
			// already in the place
			return
		}
		v := row[c]

		if v >= 0 && vis[v] == cur {
			// a cycle
			sz := 1
			w := u
			for w != v {
				sz++
				w = par[w]
			}
			cycles = append(cycles, sz)
			return
		}
		if v < 0 || vis[v] != 0 {
			return
		}
		par[v] = u
		dfs(v, cur)
	}

	for i := 0; i < m; i++ {
		if vis[i] == 0 {
			dfs(i, i+1)
		}
	}
	var sum int
	for _, x := range cycles {
		sum += x
	}
	return len(cycles) + m - good
}
