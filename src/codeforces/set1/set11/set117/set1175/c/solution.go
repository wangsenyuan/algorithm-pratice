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
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k)

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

const inf = 1 << 60

func solve(a []int, k int) int {

	ans := []int{inf, -1}

	n := len(a)
	for i := 0; i+k < n; i++ {
		dist := a[i+k] - a[i]
		if dist < ans[0] {
			ans[0] = dist
			ans[1] = a[i] + dist/2
		}
	}

	return ans[1]
}

func solve1(a []int, k int) int {
	k++
	sort.Ints(a)
	n := len(a)

	check := func(x int) int {
		// 在一个宽度为2*x的窗口内，是否有k个下标
		var l, r int
		pos := a[l] + x
		for r+1 < n && a[r+1]-pos <= x {
			r++
		}
		if r-l+1 >= k {
			return min(pos, a[r]-x)
		}
		for r+1 < n {
			d1 := a[l+1] - (pos - x)
			d2 := a[r+1] - (pos + x)
			if d1 <= d2 {
				pos += d1
			} else {
				pos += d2
			}
			for pos-a[l] > x {
				l++
			}
			for r+1 < n && a[r+1]-pos <= x {
				r++
			}
			if r-l+1 >= k {
				return pos
			}
		}
		return -inf
	}

	dist := a[n-1] - a[0]

	best := sort.Search(dist, func(w int) bool {
		return check(w) > -inf
	})

	return check(best)
}
