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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, x, y := readThreeNums(reader)
		a := readNNums(reader, n)
		res := solve(x, y, a)
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

const inf = 1 << 60

func solve(x int, y int, opponents []int) int {
	sort.Ints(opponents)

	n := len(opponents)

	// t[i] = 要胜过i时，需要的初始的score
	t := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 && b[i-1] >= opponents[i] {
			t[i] = t[i-1]
			b[i] = b[i-1] + 1
		} else {
			t[i] = opponents[i] - i
			b[i] = opponents[i] + 1
		}
	}

	var ans int

	for x < y {
		pos := search(n, func(i int) bool {
			return t[i] > x
		})
		p, m := pos, n-pos
		if x+p >= y {
			ans += y - x
			return ans
		}
		if p <= m {
			return -1
		}
		k := (y - x - p + p - m - 1) / (p - m)
		if pos < n {
			k = min(k, (t[pos]-x+p-m-1)/(p-m))
		}
		ans += k * n
		x += k * (p - m)
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
