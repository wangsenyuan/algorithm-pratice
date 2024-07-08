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
		n := readNum(reader)
		a := readNNums(reader, n)
		q := readNum(reader)
		queries := make([][]int, q)
		for i := 0; i < q; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, queries)
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

func solve(a []int, queries [][]int) [][]int {
	n := len(a)
	prev := make([]int, n+1)
	prev[0] = -1
	for i := 1; i < n; i++ {
		if a[i] != a[i-1] {
			prev[i] = i - 1
		} else {
			prev[i] = prev[i-1]
		}
	}

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		l--
		r--
		if prev[r] < l {
			ans[i] = []int{-1, -1}
		} else {
			ans[i] = []int{prev[r] + 1, r + 1}
		}
	}
	return ans
}
