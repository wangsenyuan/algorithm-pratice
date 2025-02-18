package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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
	for i := range len(s) {
		if s[i] == '\r' || s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) []int {
	n, k, q := readThreeNums(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, k)
	}
	queries := make([][]string, q)
	for i := range q {
		m := readNum(reader)
		queries[i] = make([]string, m)
		for j := range m {
			queries[i][j] = readString(reader)
		}
	}

	return solve(a, queries)
}

func solve(a [][]int, queries [][]string) []int {
	n := len(a)
	k := len(a[0])
	b := make([][]int, k)
	for i := range k {
		b[i] = make([]int, n)
	}
	for i := range n {
		for j := range k {
			b[j][i] = a[i][j]
		}
	}

	for i := range k {
		for j := 1; j < n; j++ {
			b[i][j] |= b[i][j-1]
		}
	}

	query := func(conditions []string) int {
		lo, hi := 0, n

		for _, s := range conditions {
			oi := strings.IndexByte(s, ' ')
			r, _ := strconv.Atoi(s[:oi])
			r--
			v, _ := strconv.Atoi(s[oi+3:])
			oi++
			if s[oi] == '<' {
				// region r < v
				i := sort.SearchInts(b[r], v)
				hi = min(hi, i)
			} else {
				// r > v
				i := sort.SearchInts(b[r], v+1)
				lo = max(lo, i)
			}
			if lo >= hi {
				return -1
			}
		}
		return lo + 1
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = query(cur)
	}

	return ans
}
