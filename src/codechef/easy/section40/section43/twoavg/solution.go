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
		n, m, k := readThreeNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, m)
		res := solve(k, A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve1(k int, A []int, B []int) int64 {
	if k == 1 {
		return -1
	}
	sa := sum_of(A)
	sb := sum_of(B)
	n := int64(len(A))
	m := int64(len(B))

	K := int64(k)
	// add k to A, and add 1 to B
	// (sa + x * k) / (n + x) > (sb + y) / (m + y)

	best := n + m + 2

	for x, y := int64(0), n+m+2; x <= n+m+2; x++ {
		for y >= 0 && (sa+x*K)*(m+y) > (sb+y)*(n+x) {
			y--
		}
		y++
		best = min(best, x+y)
	}

	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func solve(k int, A []int, B []int) int64 {
	if k == 1 {
		return -1
	}
	sa := sum_of(A)
	sb := sum_of(B)
	n := int64(len(A))
	m := int64(len(B))

	K := int64(k)
	// add k to A, and add 1 to B
	// (sa + x * k) / (n + x) > (sb + y) / (m + y)
	x := search(n+m+1, func(i int64) bool {
		return (sa+i*K)*m > sb*(n+i)
	})

	best := x
	for x > 0 {
		x--
		y := search(n+m+1, func(i int64) bool {
			return (sa+x*K)*(m+i) > (sb+i)*(n+x)
		})

		best = min(best, x+y)
	}

	return best
}

func search(n int64, f func(int64) bool) int64 {
	var l, r int64 = 0, n
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
func sum_of(arr []int) int64 {
	var res int64
	for _, num := range arr {
		res += int64(num)
	}

	return res
}
