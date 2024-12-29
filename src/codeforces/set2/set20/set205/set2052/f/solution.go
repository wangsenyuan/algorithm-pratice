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
		parquet := make([]string, 2)
		for i := range 2 {
			parquet[i] = readString(reader)
		}
		res := solve(parquet)
		buf.WriteString(res)
		buf.WriteByte('\n')
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
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

func solve(parquet []string) string {
	// . if bad, * for good
	n := len(parquet[0])

	get := func(r int, c int) int {
		if parquet[r][c] == '#' {
			return 0
		}
		return 1
	}

	// dp[i] == 1, 表示i代表的state为true
	dp := make([]int, 4)
	dp[0] = 1
	fp := make([]int, 4)

	update := func(state int, v int) {
		fp[state] += v
		fp[state] = min(fp[state], 2)
	}

	for i := 0; i < n; i++ {
		cur := get(1, i)<<1 | get(0, i)
		// 水平放置
		if cur&1 == 1 {
			update(cur^1, dp[1])
		}
		if cur&2 == 2 {
			update(cur^2, dp[2])
		}
		// 不放置的时候，前面必须已经放好了
		update(cur, dp[0])

		if cur == 3 {
			// 垂直放置的时候， 前面也必须放好了，或者前面也可以垂直
			update(0, dp[0]+dp[3])
		}

		copy(dp, fp)
		clear(fp)
	}
	if dp[0] == 0 {
		// 没有答案
		return "None"
	}

	if dp[0] == 1 {
		return "Unique"
	}

	return "Multiple"
}
