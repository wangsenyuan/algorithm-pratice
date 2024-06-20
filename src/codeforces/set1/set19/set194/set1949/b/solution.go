package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(x []int, y []int) int {
	sort.Ints(x)
	sort.Ints(y)

	n := len(x)

	check := func(l int) int {
		res := 1 << 30
		for i := 0; i < l; i++ {
			res = min(res, abs(x[i]-y[n-l+i]))
		}
		for i := l; i < n; i++ {
			res = min(res, abs(x[i]-y[i-l]))
		}

		return res
	}

	var best int
	for l := 0; l <= n; l++ {
		best = max(best, check(l))
	}

	return best
}

func solve1(x []int, y []int) int {
	n := len(x)
	sort.Ints(x)
	sort.Ints(y)

	check := func(expect int) bool {
		// abs(x[0] - y[i]) >= expect
		// 存在x的一个前缀，和y的相同长度的后缀匹配
		// 假设这个长度是l
		for l := 0; l <= n; l++ {
			ok := true
			for i := 0; i < l; i++ {
				if abs(x[i]-y[n-l+i]) < expect {
					ok = false
					break
				}
			}
			if ok {
				for i := l; i < n; i++ {
					if abs(x[i]-y[i-l]) < expect {
						ok = false
						break
					}
				}
			}
			if ok {
				return true
			}
		}
		return false
	}

	mx := max(abs(x[0]-y[n-1]), abs(x[n-1]-y[0]))

	l, r := 0, mx+1
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

func abs(num int) int {
	return max(num, -num)
}
