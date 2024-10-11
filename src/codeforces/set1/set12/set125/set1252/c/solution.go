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

	row := readNNums(reader, n)
	col := readNNums(reader, n)

	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 4)
	}

	res := solve(row, col, queries)

	var buf bytes.Buffer

	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(row []int, col []int, queries [][]int) []bool {
	n := len(row)

	get := func(arr []int) []int {
		dp := make([]int, n)
		// dp[i] = 左边和i奇偶性不同的最近的下标
		pos := []int{-1, -1}
		for i, x := range arr {
			x = x & 1
			dp[i] = pos[1^x]
			pos[x] = i
		}

		return dp
	}

	pr := get(row)
	pc := get(col)

	ans := make([]bool, len(queries))

	check := func(arr []int, l int, r int) bool {
		if l > r {
			l, r = r, l
		}

		return arr[r-1] < l-1
	}

	for i, cur := range queries {
		ra, ca, rb, cb := cur[0], cur[1], cur[2], cur[3]
		ans[i] = check(pr, ra, rb) && check(pc, ca, cb)
	}

	return ans
}
