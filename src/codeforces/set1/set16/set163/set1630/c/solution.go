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

const INF = 1000000007

func solve(A []int) int {
	n := len(A)

	first := make([]int, n+1)
	last := make([]int, n+1)
	for i := 0; i <= n; i++ {
		first[i] = -1
		last[i] = -1
	}
	for i, a := range A {
		if first[a] < 0 {
			first[a] = i
		}
		last[a] = i
	}

	m := n + 1
	arr := make([]int, 2*m)

	set := func(p int, v int) {
		p += m
		arr[p] = v
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		res := INF
		l += m
		r += m
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = INF
		set(i, INF)
	}
	dp[0] = 0
	set(0, 0)

	for i := 0; i < n; i++ {
		a := A[i]
		if last[a] == i {
			x := dp[first[a]] + 1
			if first[a] < last[a] {
				x++
			}

			dp[i+1] = min(dp[i+1], x)

			if first[a]+1 <= last[a]-1 {
				dp[i+1] = min(dp[i+1], get(first[a]+2, last[a]+1)+1)
			}
		}
		set(i+1, dp[i+1])
	}

	return n - dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
