package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, m)
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const inf = 1 << 30

func solve(A []int, B []int) int {
	// 如果 A[b[0]] >= A[b[1]] 不成立，=> -1
	// 或者 A[b[1]] - A[b[0]] < b[1] - b[0]
	// 然后就是每段要找出，使得它们递增的最少次数
	n := len(A)

	buf := make([]int, n+2)
	buf[0] = -inf
	copy(buf[1:], A)
	buf[n+1] = inf
	for i := 1; i <= n; i++ {
		buf[i] -= i
	}
	prev := 0

	lis := make([]int, n)
	var res int
	for i := 0; i <= len(B); i++ {
		cur := n + 1
		if i < len(B) {
			cur = B[i]
		}
		if buf[prev] > buf[cur] {
			return -1
		}
		// buf[prev] <= buf[cur]
		var p int
		for j := prev + 1; j < cur; j++ {
			if buf[prev] <= buf[j] && buf[j] <= buf[cur] {
				k := search(p, func(k int) bool {
					return lis[k] > buf[j]
				})
				lis[k] = buf[j]
				if k == p {
					p++
				}
			}
		}
		res += cur - prev - 1 - p
		prev = cur
	}
	return res
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
