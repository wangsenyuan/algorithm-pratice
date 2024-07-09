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
		n, m := readTwoNums(reader)
		s := readString(reader)[:n]
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(s, queries)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", len(x)))
			if len(x) > 0 {
				for _, r := range x {
					buf.WriteString(fmt.Sprintf("%d ", r))
				}
				buf.WriteByte('\n')
			}
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

func solve(s string, queries [][]int) [][]int {
	n := len(s)
	p := make([]int, n+1)
	for i := 0; i < n; i++ {
		p[i+1] = p[i]
		var cur int
		if s[i] == '+' {
			cur++
		} else {
			cur--
		}
		if i%2 == 0 {
			p[i+1] += cur
		} else {
			p[i+1] -= cur
		}
	}

	get := func(l int, r int) int {
		if l > r {
			return 0
		}
		if l%2 == 1 {
			return p[r] - p[l-1]
		}
		return p[l-1] - p[r]
	}

	check := func(l int, r int, m int) int {
		if (m-l+1)%2 == 1 {
			return get(l, m-1) + get(m+1, r)
		}
		return get(l, m-1) - get(m+1, r)
	}

	find_odd_seg := func(l int, r int) int {
		if l == r {
			return l
		}
		lb, rb := l, r
		for lb < rb {
			mid := (lb + rb) / 2
			lq := check(l, r, lb)
			mq := check(l, r, mid)
			rq := check(l, r, rb)
			if lq == 0 {
				return lb
			}
			if mq == 0 {
				return mid
			}
			if rq == 0 {
				return rb
			}
			if sign(lq) == sign(mq) {
				lb = mid
			} else {
				rb = mid
			}
		}
		return 0
	}

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		if get(l, r) == 0 {
			ans[i] = []int{}
			continue
		}
		var even bool
		if (r-l+1)%2 == 0 {
			even = true
			l++
		}
		pos := find_odd_seg(l, r)
		if even {
			ans[i] = []int{l - 1, pos}
		} else {
			ans[i] = []int{pos}
		}
	}
	return ans
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}
