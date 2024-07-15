package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)
	for tc > 0 {
		n := readNum(reader)
		a := readNNums(reader, n)
		m := readNum(reader)
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, queries)

		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}

		tc--
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

func solve(a []int, queries [][]int) []int {
	//显然不会往后移动

	n := len(a)
	pref := make([]int, n+1)

	for i := 1; i < n; i++ {
		pref[i+1] = pref[i]
		if i == n-1 || a[i]-a[i-1] < a[i+1]-a[i] {
			// 前面的那个是更close的
			pref[i+1]++
		} else {
			pref[i+1] += a[i] - a[i-1]
		}
	}

	suf := make([]int, n+1)
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1]
		if i == 0 || a[i+1]-a[i] < a[i]-a[i-1] {
			suf[i]++
		} else {
			suf[i] += a[i+1] - a[i]
		}
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		if l < r {
			ans[i] = suf[l-1] - suf[r-1]
		} else {
			ans[i] = pref[l] - pref[r]

		}
	}

	return ans
}
