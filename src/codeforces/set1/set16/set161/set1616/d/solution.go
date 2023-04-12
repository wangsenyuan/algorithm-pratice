package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		x := readNum(reader)
		res := solve(A, x)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(A []int, x int) int {
	n := len(A)
	for i := 0; i < n; i++ {
		A[i] -= x
	}

	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if A[i]+A[i-1] < 0 && dp[i-1] == 1 {
			dp[i] = 0
		} else if i >= 2 && A[i-2]+A[i-1]+A[i] < 0 && dp[i-2] == 1 && dp[i-1] == 1 {
			dp[i] = 0
		} else {
			dp[i] = 1
		}
	}
	var sum int
	for i := 0; i < n; i++ {
		sum += dp[i]
	}
	return sum
}

func solve1(A []int, x int) int {
	// l...r 中间，要们
	// 1）至少有一个元素没有被选
	// 2) sum(r) - sum(l - 1) < x * (r - l + 1)
	n := len(A)

	N := n + 1
	arr := make([]int64, 2*N)
	for i := 0; i < len(arr); i++ {
		arr[i] = -INF
	}

	set := func(p int, v int64) {
		p += N
		arr[p] = v
		for p > 1 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int64 {
		l += N
		r += N
		var res int64 = -INF
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

	var sum int64
	set(0, 0)
	var l int
	var seg int = 1
	for i := 1; i <= n; i++ {
		sum += int64(A[i-1])
		val := sum - int64(x)*int64(i-l)

		tmp := get(l, i-1)
		if tmp > val {
			seg++
			sum = 0
			val = 0
			l = i
		}
		set(i, val)
	}

	return n - (seg - 1)
}

type Num interface {
	int | int64
}

func max[T Num](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

const INF = 1 << 60
