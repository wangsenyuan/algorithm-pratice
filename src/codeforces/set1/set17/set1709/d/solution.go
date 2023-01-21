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

	A := readNNums(reader, m)

	q := readNum(reader)

	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 5)
	}

	res := solve(n, m, A, Q)

	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		if res[i] {
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(height int, m int, A []int, Q [][]int) []bool {

	arr := make([]int, 2*m)
	copy(arr[m:], A)

	for i := m - 1; i > 0; i-- {
		arr[i] = max(arr[2*i], arr[2*i+1])
	}

	get := func(l int, r int) int {
		var res int
		l += m
		r += m

		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	check := func(a, b []int, k int) bool {
		y1, x1 := a[0], a[1]
		x1--
		y2, x2 := b[0], b[1]
		x2--
		if x1%k != x2%k || y1%k != y2%k {
			return false
		}
		h := get(min(x1, x2), max(x1, x2)+1)

		// h >= min(y1, y2)
		y := max(y1, y2)
		// y + k * cnt <= height
		cnt := (height - y) / k
		y += cnt * k
		return y > h
	}

	ans := make([]bool, len(Q))

	for i, cur := range Q {
		ans[i] = check(cur[:2], cur[2:4], cur[4])
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
