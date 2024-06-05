package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	p := readNNums(reader, n)
	q := readNNums(reader, n)
	res := solve(p, q)
	fmt.Println(res)
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

func solve(p []int, q []int) int {
	n := len(p)
	for i := 0; i < n; i++ {
		p[i]--
		q[i]--
	}

	ps := getPosition(p)
	qs := getPosition(q)

	first := min(ps[0], qs[0])
	last := max(ps[0], qs[0])

	get := func(m int) int {
		if m <= 0 {
			return 0
		}
		return m * (m + 1) / 2
	}

	// mex = 0 的区间
	var res int
	res += get(first)
	res += get(last - first - 1)
	res += get(n - last - 1)

	for x := 1; x < n; x++ {
		l := min(ps[x], qs[x])
		r := max(ps[x], qs[x])
		// mex = x的区间， 必须包含 [first...last], 但是不能包含(l, r)
		if l < first && last < r {
			res += (first - l) * (r - last)
		} else if l > last {
			res += (first + 1) * (l - last)
		} else if r < first {
			res += (first - r) * (n - last)
		}

		first = min(first, l)
		last = max(last, r)
	}

	return res + 1
}

func getPosition(nums []int) []int {
	n := len(nums)
	pos := make([]int, n)
	for i, num := range nums {
		pos[num] = i
	}
	return pos
}
