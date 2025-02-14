package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	pairs := make([][]int, n)

	for i := 0; i < n; i++ {
		pairs[i] = readNNums(reader, 2)
	}

	res := solve(pairs)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(pairs [][]int) []int {
	n := len(pairs)
	// 转折越少越好？

	type record struct {
		lo int
		hi int
	}

	records := make([]record, n)
	records[n-1] = record{pairs[n-1][0], pairs[n-1][1]}

	for i := n - 2; i >= 0; i-- {
		l, r := records[i+1].lo, records[i+1].hi
		if max(l, pairs[i][0]) <= min(r, pairs[i][1]) {
			// 有重叠区域
			records[i] = record{max(l, pairs[i][0]), min(r, pairs[i][1])}
		} else if pairs[i][1] < l {
			records[i] = record{pairs[i][1], pairs[i][1]}
		} else {
			records[i] = record{pairs[i][0], pairs[i][0]}
		}
	}

	ans := make([]int, n)

	ans[0] = records[0].lo

	for i := 1; i < n; i++ {
		ans[i] = max(min(ans[i-1], records[i].hi), pairs[i][0])
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
