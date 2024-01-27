package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	ask := func(i, j int) int {
		fmt.Printf("? %d %d\n", i, j)
		return readNum(reader)
	}
	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n, ask)
		fmt.Printf("! %d\n", res)
	}

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

func solve(n int, ask func(int, int) int) int {
	fp := make([][]int, n)
	for i := 0; i < n; i++ {
		fp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			fp[i][j] = -1
		}
		fp[i][i+1] = i
	}

	query := func(i, j int) int {
		if i == j {
			return 0
		}
		return ask(i+1, j+1)
	}

	var dfs func(l int, r int)

	dfs = func(l int, r int) {
		if l+1 == r {
			return
		}

		mid := (l + r) / 2
		//dfs(l, mid)
		dfs(mid, r)
		//i := fp[l][mid]
		j := fp[mid][r]
		a := query(l, j-1)
		b := query(l, j)

		if a == b {
			fp[l][r] = j
		} else {
			// a < b
			dfs(l, mid)
			fp[l][r] = fp[l][mid]
		}
	}

	dfs(0, n)

	return fp[0][n] + 1
}
