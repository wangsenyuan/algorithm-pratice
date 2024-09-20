package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(a, b, k)
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

type pair struct {
	first  int
	second int
}

func solve(a []int, b []int, k int) int {
	if k == 0 {
		return 0
	}

	check := func(base int) pair {
		// 在最后一次插入前的数字，都是大于等于base的
		var cnt int
		var score int
		for i, x := range a {
			if x < base {
				continue
			}
			y := b[i]
			tmp := (x - base) / y
			cnt += tmp + 1
			score += (x + x - tmp*y) * (tmp + 1) / 2
		}

		return pair{cnt, score}
	}

	ma := slices.Max(a)

	l, r := 0, ma+1
	for l < r {
		mid := (l + r) / 2
		tmp := check(mid)
		if tmp.first <= k {
			r = mid
		} else {
			l = mid + 1
		}
	}

	res := check(r)
	k -= res.first
	// 这里的是连续的了，所以剩下的部分，可以选择r-1
	score := res.second + k*max(0, r-1)

	return score
}
