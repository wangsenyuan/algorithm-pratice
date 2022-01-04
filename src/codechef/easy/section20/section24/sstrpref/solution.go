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
		S1 := readString(reader)
		S2 := readString(reader)
		X := readString(reader)
		res := solve(S1, S2, X)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Println(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(S1, S2 string, X string) int64 {
	X = X + "#"
	z1 := zFunction(S1 + "?" + X)[len(S1)+1:]
	z2 := zFunction(S2 + "?" + X)[len(S2)+1:]
	n := len(X)

	B := NewRMQ(n)

	for i := 0; i < n; i++ {
		B.Set(i, z2[i]+i)
	}

	var ans int64
	for l := n - 1; l >= 0; l-- {
		// z1[l] + max(B[l]..., l + z1[l])
		r := B.FindMax(l, z1[l]+l+1)
		ans += int64(r - l)
	}

	return ans
}

func solve1(S1, S2 string, X string) int64 {
	X = X + "#"
	z1 := zFunction(S1 + "?" + X)[len(S1)+1:]
	z2 := zFunction(S2 + "?" + X)[len(S2)+1:]
	n := len(X)

	B := make([][]int, n)

	for i := 0; i < n; i++ {
		B[i] = make([]int, 22)
		B[i][0] = z2[i] + i
	}

	for j := 1; j < 22; j++ {
		for i := 0; i+(1<<uint(j)) <= n; i++ {
			B[i][j] = max(B[i][j-1], B[i+(1<<uint(j-1))][j-1])
		}
	}

	rmq := func(l, r int) int {
		j := lg(r - l + 1)
		return max(B[l][j], B[r-(1<<uint(j))+1][j])
	}

	var ans int64
	for l := n - 1; l >= 0; l-- {
		// z1[l] + max(B[l]..., l + z1[l])
		r := rmq(l, z1[l]+l)
		ans += int64(r - l)
	}

	return ans
}

func lg(n int) int {
	var i int
	for 1<<uint(i+1) <= n {
		i++
	}
	return i
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}

	return z
}

type RMQ struct {
	arr []int
	sz  int
}

func NewRMQ(n int) *RMQ {
	arr := make([]int, 2*n)

	return &RMQ{arr, n}
}

func (q *RMQ) Set(p int, v int) {
	p += q.sz
	arr := q.arr
	arr[p] = max(arr[p], v)
	for p > 0 {
		arr[p>>1] = max(arr[p], arr[p^1])
		p >>= 1
	}
}

func (q *RMQ) FindMax(l, r int) int {
	l += q.sz
	r += q.sz
	arr := q.arr
	var res int
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
