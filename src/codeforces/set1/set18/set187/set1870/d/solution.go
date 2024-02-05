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
		k := readNum(reader)
		res := solve(a, k)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(fmt.Sprintf("%s\n", s[1:len(s)-1]))
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

func solve(c []int, k int) []int {
	n := len(c)
	for i := n - 2; i >= 0; i-- {
		c[i] = min(c[i], c[i+1])
	}

	a := make([]int, n)

	a[0] = k / c[0]
	x := k % c[0]

	for i := 1; i < n; i++ {
		if c[i] > c[i-1] {
			a[i] = min(a[i-1], x/(c[i]-c[i-1]))
			x -= a[i] * (c[i] - c[i-1])
		} else {
			a[i] = a[i-1]
		}
	}
	return a
}
func solve1(c []int, k int) []int {
	n := len(c)
	stack := make([]Pair, n)
	var top int

	stack[top] = Pair{c[n-1], n - 1}
	top++

	for i := n - 2; i >= 0; i-- {
		if c[i] >= stack[top-1].first {
			continue
		}
		stack[top] = Pair{c[i], i}
		top++
	}

	for i, j := 0, top-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	b := make([]int, top)
	for i := 0; i < top; i++ {
		cur := stack[i]
		if i == 0 {
			b[0] = k / cur.first
			k -= b[0] * cur.first
		} else {
			prev := stack[i-1]
			tot := k + b[i-1]*prev.first
			j := search(tot/prev.first, func(j int) bool {
				return (tot-j*prev.first)/prev.first == (tot-j*prev.first)/cur.first
			})
			// j个选项还是使用前面的
			k += (b[i-1] - j) * prev.first
			b[i-1] = j
			b[i] = k / cur.first
			k -= b[i] * cur.first
		}
	}

	a := make([]int, n+1)

	for i := 0; i < top; i++ {
		a[0] += b[i]
		a[stack[i].second+1] -= b[i]
	}

	for i := 1; i < n; i++ {
		a[i] += a[i-1]
	}

	return a[:n]
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type Pair struct {
	first  int
	second int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
