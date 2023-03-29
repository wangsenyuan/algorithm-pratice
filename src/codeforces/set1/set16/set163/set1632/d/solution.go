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

	tc := 1

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(A []int) []int {
	/*
		gcd(b...) = r - l + 1

		b[l...r] 中的 1 肯定要被change
		b[l, l + 1] = 2的，也需要改变
		b[i, i + 1, i + 2] = 3的也需要改变。。。

		dp[i] 表示到i为止，最少的次数
		假设 gcd(a[j], a[j+1], ... a[i]) = i - j + 1
		// 那么这时候只要将 a[i] 变成某个数字 d > 1, 使得 gcd(...) = 1， 那么 dp[i] = dp[j-1] + 1
		//  gcd(a[j], gcd(a[j+1], ...)) = L
		//  a[j] % L = 0
		//  g[j+1] % L = 0
		以i为止的g个个数不会很多
	*/
	n := len(A)

	arr := make([]int, 2*n)

	for i := n; i < 2*n; i++ {
		arr[i] = A[i-n]
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = gcd(arr[2*i], arr[2*i+1])
	}

	get := func(l, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res = gcd(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = gcd(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	dp := make([]int, n+1)
	for l, r := 1, 1; r <= n; r++ {
		// g is gcd(l...r-1)
		for get(l-1, r) < r-l+1 {
			l++
		}
		// l <= r, gcd(A[r]) = A[r] >= 1
		x := get(l-1, r)
		if x == r-l+1 {
			dp[r] = dp[l-1] + 1
			l = r + 1
		} else {
			dp[r] = dp[r-1]
		}
	}

	return dp[1:]
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
